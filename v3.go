package tordn

import (
	"crypto/ed25519"
	"errors"
)

// V3 TOR address generator
type V3 struct {
}

// GenerateTORDomainName implements TORDomainNameGenerator interface
//
// Generate tor v3 domain name using the secretKey. If the secretKey is empty then default ed25519.GenerateKey
// function will be used otherwise ed25519.NewKeyFromSeed will be used to generate the public key.
// If provided the length of the secretKey must be 32
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
