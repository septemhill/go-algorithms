package sorting

func QuickSort(nums []int) []int {
	var larr, rarr []int

	l := len(nums)

	if l <= 1 {
		return nums
	}

	pivot := nums[0]

	for i := 1; i < l; i++ {
		if comparator(pivot, nums[i]) {
			larr = append(larr, nums[i])
		} else {
			rarr = append(rarr, nums[i])
		}
	}

	return intSliceConcat(QuickSort(larr), []int{pivot}, QuickSort(rarr))
}
