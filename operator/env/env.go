package env

import (
	libEnv "github.com/codingconcepts/env"
)

type Env struct {
	ClusterId       string `env:"CLUSTER_ID" required:"true"`
	ResourceIdLabel string `env:"RESOURCE_ID_LABEL" required:"true"`

	KafkaBrokers           string `env:"KAFKA_BROKERS" required:"true"`
	KafkaIncomingTopic     string `env:"KAFKA_INCOMING_TOPIC" required:"true"`
	KafkaStatusReplyTopic  string `env:"KAFKA_STATUS_REPLY_TOPIC" required:"true"`
	KafkaBillingReplyTopic string `env:"KAFKA_BILLING_REPLY_TOPIC" required:"true"`
	KafkaConsumerGroupId   string `env:"KAFKA_CONSUMER_GROUP_ID" required:"true"`

	HarborAdminUsername      string `env:"HARBOR_ADMIN_USERNAME" required:"true"`
	HarborAdminPassword      string `env:"HARBOR_ADMIN_PASSWORD" required:"true"`
	HarborImageRegistryHost  string `env:"HARBOR_IMAGE_REGISTRY_HOST" required:"true"`
	HarborApiVersion         string `env:"HARBOR_API_VERSION" required:"false"`
	HarborProjectStorageSize int    `env:"HARBOR_PROJECT_STORAGE_SIZE" required:"true"`
	HarborQuoteEnabled       bool   `env:"HARBOR_QUOTA_ENABLED" required:"true"`

	StorageClass       string `env:"STORAGE_CLASS" required:"true"`
	XFSStorageClass    string `env:"XFS_STORAGE_CLASS" required:"true"`
	ReconcilePeriod    string `env:"RECONCILE_PERIOD" required:"true"`
	ServiceAccountName string `env:"SERVICE_ACCOUNT_NAME" required:"true"`

	AwsAccessKeyId     string `env:"AWS_ACCESS_KEY_ID" required:"true"`
	AwsSecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY" required:"true"`

	ClusterCertIssuer         string `env:"CLUSTER_CERT_ISSUER" required:"true"`
	DefaultIngressClass       string `env:"DEFAULT_INGRESS_CLASS" required:"true"`
	WildcardDomainSuffix      string `env:"WILDCARD_DOMAIN_SUFFIX" required:"true"`
	WildcardDomainCertificate string `env:"WILDCARD_DOMAIN_CERTIFICATE" required:"true"`

	DockerSecretName        string `env:"DOCKER_SECRET_NAME" required:"true"`
	NamespaceAdminRoleName  string `env:"NAMESPACE_ADMIN_ROLE_NAME" required:"true"`
	NamespaceSvcAccountName string `env:"NAMESPACE_SVC_ACCOUNT_NAME" required:"true"`
}

func GetEnv() (*Env, error) {
	var env Env
	if err := libEnv.Set(&env); err != nil {
		return nil, err
	}
	return &env, nil
}

func Must(env *Env, err error) *Env {
	if err != nil {
		panic(err)
	}
	return env
}
