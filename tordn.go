package tordn

// TOR domain name generator interface
type TORDomainNameGenerator interface {
	// Generate tor domain name using the secret key
	GenerateTORDomainName(secretKey []byte) (publicKey []byte, privateKey []byte, onionAddress []byte, err error)
}
