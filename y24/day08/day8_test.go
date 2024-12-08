package day08

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkDay8Part1(b *testing.B) {
	start := time.Now()
	Day8Part1()
	fmt.Printf("Part 1: %v\n", time.Since(start))
}

func BenchmarkDay8Part2(b *testing.B) {
	start := time.Now()
	Day8Part2()
	fmt.Printf("Part 2: %v\n", time.Since(start))
}
