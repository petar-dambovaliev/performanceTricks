package strengthReduction

import "testing"

var arr = []uint64{1, 123, 432432434324324324, 123232123, 123123, 928319284371847328}

func BenchmarkDigitCount(b *testing.B) {
	for _, u := range arr {
		for i:=0;i < b.N;i++ {
			DigitCount(u)
		}
	}
}

func BenchmarkStrengthReductionDigitCount(b *testing.B) {
	for _, u := range arr {
		for i:=0;i < b.N;i++ {
			StrengthReductionDigitCount(u)
		}
	}
}
