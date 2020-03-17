package tordn

import (
	"crypto/ed25519"
	"io"
)

// V3 TOR address generator
type V3 struct {
}

// GenerateTORDomainName implements TORDomainNameGenerator interface
//
// Generate tor v3 domain name using the secretKey.
func (g *V3) GenerateTORDomainName(secretKey io.Reader) (publicKey []byte, privateKey []byte, onionAddress []byte, err error) {
	// Generate key pair
	publicKey, privateKey, err = ed25519.GenerateKey(secretKey)
	if err != nil {
		return nil, nil, nil, err
	}

	onionAddress = MakeV3OnionAddressWithExtension(publicKey)
	return publicKey, privateKey, onionAddress, nil
}
