package tordn

import (
	"bytes"
	"io"
	"testing"
)

func BenchmarkV3_GenerateTORDomainName(b *testing.B) {
	b.ReportAllocs()
	g := &V3{}
	for x := 0; x < b.N; x++ {
		_, _, _, _ = g.GenerateTORDomainName(nil)
	}
}

var torDomainList = []struct {
	secretKey            []byte
	expectedOnionAddress string
}{
	{
		[]byte("eipuopkyhrisvmrlbghubnlxunzkwoij"),
		"fsf7sfk2ebq5uj7ekd3m75vtihy7dg7ut7fhholibye4qjb7uu5px4yd.onion",
	},
}

func TestV3_GenerateTORDomainName(t *testing.T) {
	v3 := &V3{}
	for _, d := range torDomainList {
		_, _, foundOnionAddress, _ := v3.GenerateTORDomainName(bytes.NewReader(d.secretKey))
		if d.expectedOnionAddress != string(foundOnionAddress) {
			t.Errorf("Invalid onion address. Expected: [%v] Found: [%v]", d.expectedOnionAddress, foundOnionAddress)
		}
	}

	{
		publicKey, privateKey, foundOnionAddress, _ := v3.GenerateTORDomainName(nil)
		if len(foundOnionAddress) != 62 {
			t.Errorf("Expected onion address length does not match")
		}

		if len(publicKey) != 32 {
			t.Errorf("Expected publicKey length does not match")
		}

		if len(privateKey) != 64 {
			t.Errorf("Expected privateKey length does not match")
		}
	}

	{
		_, _, _, err := v3.GenerateTORDomainName(bytes.NewReader([]byte("1")))
		if err != io.ErrUnexpectedEOF {
			t.Errorf("Unexpected error %v", err)
		}
	}
}
