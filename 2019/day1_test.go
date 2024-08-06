package _2019

import "testing"

func BenchmarkDay1Part1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Day1Part1()
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Day1Part2()
	}
}
