package main

import (
	"fmt"
	"sort"
)

func main() {
	result := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println("expected 2.00, got: ", result)

	result = findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println("expected 2.50, got: ", result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)

	length := len(merged)
	if length == 0 {
		return 0
	} else if length%2 == 0 {
		return float64(merged[(length/2)-1]+merged[length/2]) / 2
	} else {
		return float64(merged[length/2])
	}
}
