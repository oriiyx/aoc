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

	fmt.Printf("Day 1 Part 1 fuel sum: %v \n", fuelReq)
}

func calculateFuelRequirement(mass int) int {
	return (mass / 3) - 2
}

func Part2() {
	requestURL := "https://adventofcode.com/2019/day/1/input"
	lines := data.GetInputArray(requestURL)

	fuelReq := 0

	for _, massString := range lines {
		absoluteFuelMass := 0
		mass, _ := strconv.Atoi(massString)
		fuelReq += calculateAbsoluteFuelRequirement(mass, absoluteFuelMass)
	}

	fmt.Printf("Day 1 Part 2 fuel sum: %v \n", fuelReq)
}

func calculateAbsoluteFuelRequirement(mass int, absoluteMass int) int {
	requirement := calculateFuelRequirement(mass)

	if requirement > 0 {
		absoluteMass += requirement
		absoluteMass = calculateAbsoluteFuelRequirement(requirement, absoluteMass)
	}

	return absoluteMass
}
