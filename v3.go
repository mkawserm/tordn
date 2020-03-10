package tordn

import (
	"crypto/ed25519"
	"errors"
)

// Tor v3 onion address generator
type V3 struct {
}

// Generate tor v3 domain name using the secret key if secretKey is empty using default ed25519
// key generation method. The length of the secretKey must be 32
func (g *V3) GenerateTORDomainName(secretKey []byte) (publicKey []byte, privateKey []byte, onionAddress []byte, err error) {
	// Generate key pair
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

	onionAddress = MakeV3OnionAddressWithExtension(publicKey)
	return publicKey, privateKey, onionAddress, nil
}
