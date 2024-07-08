package _2019

import (
	"fmt"
	"strconv"

	"github.com/oriiyx/aoc/internal/data"
)

func Part1() {
	requestURL := "https://adventofcode.com/2019/day/1/input"
	lines := data.GetInputArray(requestURL)

	fuelReq := 0
	for _, massString := range lines {
		mass, _ := strconv.Atoi(massString)
		fuelReq += calculateFuelRequirement(mass)
	}

	fmt.Printf("Day 1 fuel sum: %v", fuelReq)
}

func calculateFuelRequirement(mass int) int {
	return (mass / 3) - 2
}

func Part2() {

}
