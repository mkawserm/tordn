package tordn

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base32"
	"golang.org/x/crypto/sha3"
)

type TORV3DomainNameGenerator struct {
}

func (t *TORV3DomainNameGenerator) GenerateTORDomainName() (TORPublicKey, TORPrivateKey, TORDomainName, error) {

	// Generate key pair
	publicKey, privateKey, err := ed25519.GenerateKey(nil)

	if err != nil {
		return nil, nil, "", err
	}

	// checksum = H(".onion checksum" || publicKey || version)
	var checksumBytes bytes.Buffer
	checksumBytes.Write([]byte(".onion checksum"))
	checksumBytes.Write(publicKey)
	checksumBytes.Write([]byte{0x03})
	checksum := sha3.Sum256(checksumBytes.Bytes())

	// onion_address = base32(publicKey || checksum || version)
	var onionAddressBytes bytes.Buffer
	onionAddressBytes.Write(publicKey)
	onionAddressBytes.Write(checksum[:2])
	onionAddressBytes.Write([]byte{0x03})
	onionAddress := base32.StdEncoding.EncodeToString(onionAddressBytes.Bytes())

	return TORPublicKey(publicKey), TORPrivateKey(privateKey), TORDomainName(onionAddress), nil
}
