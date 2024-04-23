package wg

import (
	"fmt"

	"github.com/seancfoley/ipaddress-go/ipaddr"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func Ptr[T any](t T) *T {
	return &t
}

func GenerateWgKeys() ([]byte, []byte, error) {
	key, err := wgtypes.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	return []byte(key.PublicKey().String()), []byte(key.String()), nil
}

func GeneratePublicKey(privateKey string) ([]byte, error) {
	key, err := wgtypes.ParseKey(privateKey)
	if err != nil {
		return nil, err
	}

	return []byte(key.PublicKey().String()), nil
}

func GetRemoteDeviceIp(deviceOffset int64, IpBase string) ([]byte, error) {
	deviceRange := ipaddr.NewIPAddressString(fmt.Sprintf("%s/16", IpBase))

	if address, addressError := deviceRange.ToAddress(); addressError == nil {
		increment := address.Increment(deviceOffset)
		return []byte(ipaddr.NewIPAddressString(increment.GetNetIP().String()).String()), nil
	} else {
		return nil, addressError
	}
}
