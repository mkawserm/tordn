package tordn

import "testing"

func BenchmarkTORV3DomainNameGenerator_GenerateTORDomainName(b *testing.B) {
	b.ReportAllocs()
	g := &TORV3DomainNameGenerator{}
	for x := 0; x < b.N; x++ {
		_, _, _, _ = g.GenerateTORDomainName()
	}
}
