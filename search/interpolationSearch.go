package search

func InterpolatioSearch(nums []int, target int) int {
	l := len(nums)
	start, end := 0, l-1

	for (nums[end] != nums[start]) && (target >= nums[start]) && (target <= nums[end]) {
		mid := start + ((target - nums[start]) * (end - start) / (nums[end] - nums[start]))

		if nums[mid] < target {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}

	if target == nums[start] {
		return start
	}

	return -1
}
