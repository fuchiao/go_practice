package main

import (
	"testing"
)

func BenchmarkStringConcat1(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concat1("zack")
	}
}
func BenchmarkStringConcat2(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concat2("zack")
	}
}
func BenchmarkStringConcat3(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concat3("zack")
	}
}
func BenchmarkStringConcat4(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		concat4("zack")
	}
}
