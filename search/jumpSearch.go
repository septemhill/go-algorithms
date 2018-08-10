package search

import (
	"math"
)

func JumpSearch(nums []int, target int) int {
	l := len(nums)
	jumpSize := int(math.Floor(math.Sqrt(float64(l))))
	start, end := 0, 0

	for start < l && nums[start] <= target {
		end = int(math.Min(float64(l-1), float64(start+jumpSize)))

		if nums[start] <= target && nums[end] >= target {
			break
		}

		start += jumpSize
	}

	end = int(math.Min(float64(l-1), float64(end)))

	for i := start; i <= end && nums[start] <= target; i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
