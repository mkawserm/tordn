package tordn

type TORPublicKey []byte
type TORPrivateKey []byte
type TORDomainName []byte

type TORDomainNameGenerator interface {
	GenerateTORDomainName() (TORPublicKey, TORPrivateKey, TORDomainName, error)
}
