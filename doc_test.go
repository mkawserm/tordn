package tordn_test

import (
	"bytes"
	"fmt"
	"github.com/mkawserm/tordn"
)

// Generate v3 tor domain name
func ExampleV3_GenerateTORDomainName() {

	v3domainName := &tordn.V3{}

	// generate random v3 tor domain name
	publicKey, privateKey, onionAddress, err := v3domainName.GenerateTORDomainName(nil)
	if err == nil {
		fmt.Printf("Public Key:")
		fmt.Println(publicKey)

		fmt.Printf("Private Key:")
		fmt.Println(privateKey)

		fmt.Printf("Onion Address:")
		fmt.Println(string(onionAddress))
	}

	secretKey := []byte("eipuopkyhrisvmrlbghubnlxunzkwoij")
	publicKey1, privateKey1, onionAddress1, err1 := v3domainName.GenerateTORDomainName(bytes.NewReader(secretKey))
	if err1 == nil {
		fmt.Printf("Public Key:")
		fmt.Println(publicKey1)

		fmt.Printf("Private Key:")
		fmt.Println(privateKey1)

		fmt.Printf("Onion Address:")
		fmt.Println(string(onionAddress1))
	}
}
