package tordn

import (
	"bytes"
	"encoding/base32"
	"golang.org/x/crypto/sha3"
	"strings"
)

// Apply tor v3 onion address generation algorithm
func PublicKeyToV3Address(publicKey []byte) []byte {
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
	return onionAddressBytes.Bytes()
}

// Make tor v3 onion address with .onion extension
func MakeV3OnionAddressWithExtension(publicKey []byte) []byte {
	onionAddress := strings.ToLower(base32.StdEncoding.EncodeToString(PublicKeyToV3Address(publicKey)))
	onionAddress = onionAddress + ".onion"
	return []byte(onionAddress)
}
