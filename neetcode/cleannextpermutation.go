package main

func main() {
	// nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{3, 2, 1})
	// nextPermutation([]int{1, 1, 5})
	// nextPermutation([]int{1, 3, 6, 8, 7, 2})
}

func nextPermutation(nums []int) {
	numsLen := len(nums) - 1
	peakIdx := 0

	for i := numsLen - 1; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			peakIdx = i + 1
			break
		}
	}

	if peakIdx != 0 {
		for i := numsLen; i >= peakIdx; i-- {
			if nums[peakIdx-1] < nums[i] {
				nums[i], nums[peakIdx-1] = nums[peakIdx-1], nums[i]
				break
			}
		}
	}

	for i := len(nums) - 1; i > peakIdx; i, peakIdx = i-1, peakIdx+1 {
		nums[i], nums[peakIdx] = nums[peakIdx], nums[i]
	}
}
