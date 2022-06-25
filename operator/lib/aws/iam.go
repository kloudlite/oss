package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
)

type iamObj struct {
	cli *iam.IAM
}

func (i *iamObj) CreateUser(username string) error {
	_, err := i.getUser(username)
	if err != nil {
		if err2, ok := err.(awserr.Error); ok && err2.Code() == iam.ErrCodeNoSuchEntityException {
			_, err3 := i.cli.CreateUser(
				&iam.CreateUserInput{
					Path:                nil,
					PermissionsBoundary: nil,
					Tags:                nil,
					UserName:            &username,
				},
			)
			if err3 != nil {
				return err3
			}
		}
	}
	return nil
}

func (i *iamObj) DeleteUser(username string) error {
	if err := i.deleteAccessKey(username); err != nil {
		return err
	}
	_, err := i.cli.DeleteUser(
		&iam.DeleteUserInput{
			UserName: &username,
		},
	)
	if err != nil {
		if aErr, ok := err.(awserr.Error); ok && aErr.Code() == iam.ErrCodeNoSuchEntityException {
			return nil
		}
		return err
	}
	return nil
}

func (i *iamObj) CreateAccessKey(username string) (string, string, error) {
	out, err := i.cli.CreateAccessKey(
		&iam.CreateAccessKeyInput{
			UserName: aws.String(username),
		},
	)
	if err != nil {
		return "", "", err
	}
	return *out.AccessKey.AccessKeyId, *out.AccessKey.SecretAccessKey, nil
}

func (i *iamObj) DeleteAccessKey(username string) error {
	return i.deleteAccessKey(username)
}

func (i *iamObj) deleteAccessKey(username string) error {
	keysRef, err := i.cli.ListAccessKeys(
		&iam.ListAccessKeysInput{
			UserName: &username,
		},
	)
	if err != nil {
		return nil
	}

	for _, keyMeta := range keysRef.AccessKeyMetadata {
		if _, err := i.cli.DeleteAccessKey(
			&iam.DeleteAccessKeyInput{
				AccessKeyId: keyMeta.AccessKeyId,
				UserName:    &username,
			},
		); err != nil {
			return err
		}
	}
	return nil
}

func (i *iamObj) getUser(username string) (*iam.GetUserOutput, error) {
	return i.cli.GetUser(&iam.GetUserInput{UserName: &username})
}

func (i *iamObj) GetUser(username string) (*iam.GetUserOutput, error) {
	return i.getUser(username)
}

func NewIAMClient() (*iamObj, error) {
	s, err := newSession()
	if err != nil {
		return nil, err
	}
	iamCli := iam.New(s)
	return &iamObj{cli: iamCli}, nil
}
