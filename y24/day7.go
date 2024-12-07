package y24

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/oriiyx/aoc/internal/data"
)

type Operation func(a, b int) int

func Day7Part1() {
	lines, _ := data.ReadLines("day7.txt")
	count := 0

	for _, line := range lines {
		result, vals := parseLine(line)
		possible := calculate(result, vals)
		if possible {
			count += result
		}
	}

	fmt.Println(count)
}

func Day7Part2() {
	lines, _ := data.ReadLines("day7.txt")
	count := 0

	for _, line := range lines {
		result, vals := parseLine(line)
		possible := calculate2(result, vals)
		if possible {
			count += result
		}
	}

	fmt.Println(count)
}

func calculate(result int, vals []int) bool {
	operations := []Operation{
		func(a, b int) int { return a * b }, // Multiplication
		func(a, b int) int { return a + b }, // Addition
	}

	// Number of operations needed is len(vals) - 1
	numOps := len(vals) - 1
	// Total number of possible combinations is 2^numOps
	totalCombinations := 1 << numOps

	// Try each possible combination of operations
	for i := 0; i < totalCombinations; i++ {
		// Start with the first number
		currentResult := vals[0]

		// Apply operations left to right
		for j := 0; j < numOps; j++ {
			// Use bit j of i to determine which operation to use
			opIndex := (i >> j) & 1
			// Apply the operation with the next number
			currentResult = operations[opIndex](currentResult, vals[j+1])
		}

		if currentResult == result {
			return true
		}
	}

	return false
}

func calculate2(result int, vals []int) bool {
	operations := []Operation{
		func(a, b int) int { return a * b }, // Multiplication
		func(a, b int) int { return a + b }, // Addition
		func(a, b int) int {
			// Concatenation: convert both numbers to string, join them, convert back to int
			combined := fmt.Sprintf("%d%d", a, b)
			num, _ := strconv.Atoi(combined)
			return num
		},
	}

	// Number of operations needed is len(vals) - 1
	numOps := len(vals) - 1
	// Total number of possible combinations is 3^numOps since we now have 3 operations
	totalCombinations := int(math.Pow(3, float64(numOps)))

	// Try each possible combination of operations
	for i := 0; i < totalCombinations; i++ {
		// Start with the first number
		currentResult := vals[0]

		// Convert i to base-3 to try all combinations of three operations
		temp := i
		for j := 0; j < numOps; j++ {
			// Get operation index (0, 1, or 2) for this position
			opIndex := temp % 3
			temp = temp / 3

			// Apply the operation with the next number
			currentResult = operations[opIndex](currentResult, vals[j+1])
		}

		if currentResult == result {
			return true
		}
	}

	return false
}

func parseLine(line string) (int, []int) {
	s := strings.Split(line, ":")
	r := s[0]
	rInt, _ := strconv.Atoi(r)
	strVals := strings.Split(strings.TrimSpace(s[1]), " ")
	vals := make([]int, len(strVals))
	for i, str := range strVals {
		val, err := strconv.Atoi(str)
		if err != nil {
			// Handle error if needed
			continue
		}
		vals[i] = val
	}
	return rInt, vals
}
