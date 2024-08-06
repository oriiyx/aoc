package _2019

import (
	"fmt"
	"strings"

	"github.com/oriiyx/aoc/internal/data"
)

func Day2Part1() {
	requestURL := "https://adventofcode.com/2019/day/2/input"
	lines := data.GetInputArray(requestURL)
	lines = []string{"1,9,10,3,2,3,11,0,99,30,40,50"}

	for _, line := range lines {
		values := strings.Split(line, ",")
		length := len(values)
		rows := length / 4

		for valueKey, value := range values {
			if valueKey == 0 {
				switch value {
				case "1":
					fmt.Println("1 strat")
				case "2":
					fmt.Println("2 strat")
				case "99":
					fmt.Println("99 strat")
				default:
					fmt.Println("wrong")
				}
			}
		}
	}
}
