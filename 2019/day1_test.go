package _2019

import "testing"

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1()
	}
}
