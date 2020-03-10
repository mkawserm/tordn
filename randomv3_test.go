package tordn

import (
	"testing"
)

func BenchmarkRandomV3_GenerateTORDomainName(b *testing.B) {
	b.ReportAllocs()
	g := &RandomV3{}
	for x := 0; x < b.N; x++ {
		_, _, _, _ = g.GenerateTORDomainName()
	}
}
