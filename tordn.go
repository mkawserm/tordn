package tordn

// TOR domain name generator basic interface
type TORDomainNameGenerator interface {

	// Generate tor domain name using the secret key.
	//
	// Implementations must handle nil secretKey
	GenerateTORDomainName(secretKey []byte) (publicKey []byte, privateKey []byte, onionAddress []byte, err error)
}
