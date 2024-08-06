package main

import "fmt"

func main() {
	// result := twoSum([]int{2, 7, 11, 15}, 9)
	// fmt.Println(result)

	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, v := range nums {
		deducted := target - v

		index, ok := hashMap[deducted]
		if ok {
			return []int{index, i}
		}

		hashMap[v] = i
	}

	return []int{}
}
