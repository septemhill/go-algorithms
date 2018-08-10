package search

func BinarySearch(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		idx := (start + end) / 2
		if nums[idx] == target {
			return idx
		} else if nums[idx] > target {
			end = idx - 1
		} else {
			start = idx + 1
		}
	}

	return -1
}
