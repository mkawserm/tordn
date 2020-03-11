package tordn

import "testing"
import "crypto/ed25519"

var onionAddressDataList = []struct {
	torPublicKey         []byte
	expectedOnionAddress string
}{
	{
		[]byte{
			170, 91, 217, 244,
			203, 101, 56, 132,
			151, 148, 178, 245,
			12, 248, 229, 231,
			89, 91, 213, 161,
			108, 44, 147, 5,
			151, 160, 248, 61,
			70, 162, 211, 20},
		"vjn5t5glmu4ijf4uwl2qz6hf45mvxvnbnqwjgbmxud4d2rvc2mkkaead.onion",
	},
	{
		[]byte{
			197, 201, 182, 188, 206,
			249, 161, 40, 44, 70, 165,
			113, 210, 61, 44, 247, 60,
			149, 107, 43, 214, 133, 93,
			120, 10, 34, 210, 46, 104,
			98, 162, 174},
		"yxe3npgo7gqsqlcguvy5epjm646jk2zl22cv26akeljc42dcukxidsid.onion",
	},
}

func TestMakeV3OnionAddressWithExtension(t *testing.T) {

	for _, v := range onionAddressDataList {
		foundOnionAddress := string(MakeV3OnionAddressWithExtension(v.torPublicKey))
		if v.expectedOnionAddress != foundOnionAddress {
			t.Errorf("Invalid onion address. Expected: [%v] Found: [%v]", v.expectedOnionAddress, foundOnionAddress)
		}
	}
}

func BenchmarkPublicKeyToV3Address(b *testing.B) {
	b.ReportAllocs()
	// Generate key pair
	publicKey, _, _ := ed25519.GenerateKey(nil)
	for x := 0; x < b.N; x++ {
		PublicKeyToV3Address(publicKey)
	}
}
