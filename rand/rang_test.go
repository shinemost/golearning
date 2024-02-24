package rand

import (
	"math/rand"
	v2 "math/rand/v2"
	"testing"
)

const MAX = 1e9

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(MAX)
	}
}

func BenchmarkRand2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v2.IntN(MAX)
	}
}
