package tordn

type TORPublicKey []byte
type TORPrivateKey []byte
type TORDomainName string

type TORDomainNameGenerator interface {
	GenerateTORDomainName() (TORPublicKey, TORPrivateKey, TORDomainName, error)
}
