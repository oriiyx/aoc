package main

func main() {
	nextPermutationLegacy([]int{1, 2, 3})
	// nextPermutationLegacy([]int{1, 3, 6, 8, 7, 2})
}

func nextPermutationLegacy(nums []int) {
	numsLen := len(nums) - 1
	peakIndex := 0
	currHighestNum := 0
	firstLargerThenLeftPivot := 0
	leftPivot := 0
	leftPivotIndex := 0

	for i := 0; i < len(nums); i++ {
		leftPivotIndex = numsLen - i
		leftPivot = nums[leftPivotIndex]
		if leftPivot < currHighestNum {
			peakIndex = leftPivotIndex + 1
			break
		} else {
			currHighestNum = nums[leftPivotIndex]
		}
	}

	for i := numsLen; i > peakIndex-1; i-- {
		if leftPivot < nums[i] {
			firstLargerThenLeftPivot = i
			break
		}
	}

	nums = append(nums, leftPivot)
	nums[leftPivotIndex] = nums[firstLargerThenLeftPivot]
	nums = removeIndex(nums, firstLargerThenLeftPivot)

	for i := peakIndex + 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
