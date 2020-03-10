package tordn

import (
	"testing"
)

func BenchmarkV3_GenerateTORDomainName(b *testing.B) {
	b.ReportAllocs()
	g := &V3{}
	for x := 0; x < b.N; x++ {
		_, _, _, _ = g.GenerateTORDomainName(nil)
	}
}
