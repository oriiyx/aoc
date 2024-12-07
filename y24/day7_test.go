package y24

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkDay7Part1(b *testing.B) {
	start := time.Now()
	Day7Part1()
	fmt.Printf("Part 1: %v\n", time.Since(start))
}

func BenchmarkDay7Part2(b *testing.B) {
	start := time.Now()
	Day7Part2()
	fmt.Printf("Part 2: %v\n", time.Since(start))
}
