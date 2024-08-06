package main

import (
	"fmt"
)

func main() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println("expected 49... got: ", result)

	result = maxArea([]int{1, 1})
	fmt.Println("expected 1... got: ", result)
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		minHeight := height[left]
		if height[right] < minHeight {
			minHeight = height[right]
		}
		currentArea := minHeight * (right - left)
		if currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
