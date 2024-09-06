package _2019

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oriiyx/aoc/internal/data"
)

func Day2Part1() {
	requestURL := "https://adventofcode.com/2019/day/2/input"
	lines := data.GetInputArray(requestURL)
	// lines = []string{"1,9,10,3,2,3,11,0,99,1,9,10,3,2,3,11,0"}

	for _, line := range lines {
		stringArr := strings.Split(line, ",")
		var values []int

		for _, str := range stringArr {
			intVal, err := strconv.Atoi(str)
			if err == nil {
				values = append(values, intVal)
			}
		}

		values[1] = 12
		values[2] = 2
	L:
		for i := 0; i < len(values); i++ {
			switch values[i] {
			case 1:
				handleAddition(values, i)
				i = i + 3
			case 2:
				handleMultiplication(values, i)
				i = i + 3
			case 99:
				break L
			default:
				fmt.Println(fmt.Sprintf("default: %v\n", i))
			}
		}

		fmt.Println(values)
	}
}

func handleAddition(values []int, i int) {
	frstIdx := values[i+1]
	scndIdx := values[i+2]
	target := values[i+3]

	values[target] = values[frstIdx] + values[scndIdx]
}

func handleMultiplication(values []int, i int) {
	frstIdx := values[i+1]
	scndIdx := values[i+2]
	target := values[i+3]

	values[target] = values[frstIdx] * values[scndIdx]
}
