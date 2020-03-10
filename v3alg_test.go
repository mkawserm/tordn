package tordn

import "testing"
import "crypto/ed25519"

func TestMakeV3OnionAddressWithExtension(t *testing.T) {
	torPublicKey := []byte{
		170, 91, 217, 244,
		203, 101, 56, 132,
		151, 148, 178, 245,
		12, 248, 229, 231,
		89, 91, 213, 161,
		108, 44, 147, 5,
		151, 160, 248, 61,
		70, 162, 211, 20,
	}
	expectedOnionAddress := "vjn5t5glmu4ijf4uwl2qz6hf45mvxvnbnqwjgbmxud4d2rvc2mkkaead.onion"
	foundOnionAddress := string(MakeV3OnionAddressWithExtension(torPublicKey))

	if expectedOnionAddress != foundOnionAddress {
		t.Errorf("Invalid onion address. Expected: [%v] Found: [%v]", expectedOnionAddress, foundOnionAddress)
	}
}

func TestMakeV3OnionAddressWithExtension2(t *testing.T) {
	torPublicKey := []byte{
		197, 201, 182, 188, 206,
		249, 161, 40, 44, 70, 165,
		113, 210, 61, 44, 247, 60,
		149, 107, 43, 214, 133, 93,
		120, 10, 34, 210, 46, 104,
		98, 162, 174}

	expectedOnionAddress := "yxe3npgo7gqsqlcguvy5epjm646jk2zl22cv26akeljc42dcukxidsid.onion"
	foundOnionAddress := string(MakeV3OnionAddressWithExtension(torPublicKey))

	if expectedOnionAddress != foundOnionAddress {
		t.Errorf("Invalid onion address. Expected: [%v] Found: [%v]", expectedOnionAddress, foundOnionAddress)
	}
}

func TestMakeV3OnionAddressWithExtension3(t *testing.T) {
	torPublicKey := []byte{
		197, 201, 182, 188, 206,
		249, 161, 40, 44, 70, 165,
		113, 210, 61, 44, 247, 60,
		149, 107, 43, 214, 133, 93,
		120, 10, 34, 210, 46, 104,
		98, 162, 174}

	expectedOnionAddress := "vjn5t5glmu4ijf4uwl2qz6hf45mvxvnbnqwjgbmxud4d2rvc2mkkaead.onion"
	foundOnionAddress := string(MakeV3OnionAddressWithExtension(torPublicKey))

	if expectedOnionAddress == foundOnionAddress {
		t.Errorf("Expected: [%v] Found: [%v] can not be equal", expectedOnionAddress, foundOnionAddress)
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
