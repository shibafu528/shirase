package shirase

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// PEM形式でシリアライズされたRSA鍵ペア
type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

// NewKeyPair で生成するRSA鍵のビット長
const rsaKeySize = 2048

// 新しいRSA鍵ペアを生成
func NewKeyPair() (KeyPair, error) {
	key, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return KeyPair{}, fmt.Errorf("[shirase.NewKeyPair] failed to generate key: %w", err)
	}

	pair := KeyPair{
		PrivateKey: pem.EncodeToMemory(&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		}),
		PublicKey: pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(key.Public().(*rsa.PublicKey)),
		}),
	}

	return pair, nil
}
