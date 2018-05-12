package main

import (
	"testing"
)

func BenchmarkFib20(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		fib(20)
	}
}
