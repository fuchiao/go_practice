package main

import (
	"testing"
)

func BenchmarkPopcnt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		popcnt(uint64(n))
	}
}
