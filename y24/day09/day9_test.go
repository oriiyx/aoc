package day09

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkDay9Part1(b *testing.B) {
	start := time.Now()
	Day9Part1()
	fmt.Printf("Part 1: %v\n", time.Since(start))
}

func BenchmarkDay9Part2(b *testing.B) {
	start := time.Now()
	Day9Part2()
	fmt.Printf("Part 2: %v\n", time.Since(start))
}
