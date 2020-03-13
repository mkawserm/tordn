package tordn

// TORDomainNameGenerator basic interface
type TORDomainNameGenerator interface {

	// GenerateTORDomainName generates tor domain using the secret key.
	//
	// Implementations must handle nil secretKey
	GenerateTORDomainName(secretKey []byte) (publicKey []byte, privateKey []byte, onionAddress []byte, err error)
}
