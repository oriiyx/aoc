package _2019

import (
	"fmt"
	"math"
	"strconv"

	"github.com/oriiyx/aoc/internal/data"
)

func Part1() {
	requestURL := "https://adventofcode.com/2019/day/1/input"
	lines := data.GetInputArray(requestURL)

	fuelReq := 0
	for _, massString := range lines {
		mass, _ := strconv.Atoi(massString)
		fuelReq = fuelReq + calculateFuelRequirement(mass)
	}

	fmt.Println(fuelReq)
}

func calculateFuelRequirement(mass int) int {
	fuelReq := math.Floor(float64(mass/3)) - 2
	return int(fuelReq)
}
