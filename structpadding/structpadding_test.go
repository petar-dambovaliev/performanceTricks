package structpadding

import (
	"testing"
)

func BenchmarkNoPad(b *testing.B) {
	myatomic := &NoPad{}
	b.ResetTimer()
	for i:=0;i < b.N; i++ {
		go myatomic.IncreaseAllEles()
	}
}

func BenchmarkPad(b *testing.B) {
	myatomic := &Pad{}
	b.ResetTimer()
	for i:=0;i < b.N; i++ {
		go myatomic.IncreaseAllEles()
	}
}