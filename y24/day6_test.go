package y24

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkDay6Part1(b *testing.B) {
	start := time.Now()
	Day6Part1()
	fmt.Printf("Part 1: %v\n", time.Since(start))
}

func BenchmarkDay6Part2(b *testing.B) {
	start := time.Now()
	Day6Part2()
	fmt.Printf("Part 2: %v\n", time.Since(start))
}
