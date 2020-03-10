package tordn

import (
	"crypto/ed25519"
)

type RandomV3 struct {
}

func (g *RandomV3) GenerateTORDomainName() (TORPublicKey, TORPrivateKey, TORDomainName, error) {
	// Generate key pair
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, nil, nil, err
	}

	onionAddress := MakeV3OnionAddressWithExtension(publicKey)
	return TORPublicKey(publicKey), TORPrivateKey(privateKey), onionAddress, nil
}
