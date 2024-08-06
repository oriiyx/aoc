package _2019

import (
	"fmt"
	"strconv"

	"github.com/oriiyx/aoc/internal/data"
)

func Day1Part1() {
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

func Day1Part2() {
	requestURL := "https://adventofcode.com/2019/day/1/input"
	lines := data.GetInputArray(requestURL)

	fuelReq := 0

	for _, massString := range lines {
		mass, _ := strconv.Atoi(massString)
		fuelReq += calculateTotalFuelRequirement(mass)
	}

	fmt.Printf("Day 1 Part 2 fuel sum: %v \n", fuelReq)
}

func calculateTotalFuelRequirement(mass int) int {
	totalFuel := 0

	for {
		fuel := calculateFuelRequirement(mass)

		if fuel <= 0 {
			break
		}

		totalFuel += fuel
		mass = fuel
	}

	return totalFuel
}
