package secrets

import (
	"crypto/rsa"

	"github.com/rancher/go-rancher/client"
)

type Secret struct {
	client.Resource
	SecretName          string `json:"name"`
	Backend             string `json:"backend"`
	KeyName             string `json:"keyName"`
	CipherText          string `json:"cipherText"`
	ClearText           string `json:"clearText"`
	RewrapText          string `json:"rewrapText"`
	RewrapKey           string `json:"rewrapKey"`
	HashAlgorithm       string `json:"hashAlgorithm"`
	EncryptionAlgorithm string `json:"encryptionAglorigthm"`
}

type rsaPublicKey struct {
	*rsa.PublicKey
}

type encryptedData struct {
	EncryptionAlgorithm string
	EncryptedText       string
	HashAlgorithm       string
}
