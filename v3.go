package tordn

import (
	"crypto/ed25519"
	"errors"
)

type V3 struct {
}

func (g *V3) GenerateTORDomainName(secretKey []byte) ([]byte, []byte, []byte, error) {
	// Generate key pair
	var publicKey, privateKey []byte
	var err error

	if secretKey == nil {
		publicKey, privateKey, err = ed25519.GenerateKey(nil)
	} else {
		if len(secretKey) != 32 {
			return nil, nil, nil, errors.New("secretKey length must be 32")
		}

		privateKey = ed25519.NewKeyFromSeed(secretKey)
		publicKey = make([]byte, 32)
		copy(publicKey, privateKey[32:])
	}

	if err != nil {
		return nil, nil, nil, err
	}

	onionAddress := MakeV3OnionAddressWithExtension(publicKey)
	return publicKey, privateKey, onionAddress, nil
}
