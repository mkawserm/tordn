package tordn

type TORDomainNameGenerator interface {
	GenerateTORDomainName(secretKey []byte) ([]byte, []byte, []byte, error)
}
