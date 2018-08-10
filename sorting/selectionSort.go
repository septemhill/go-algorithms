package sorting

func SelectionSort(nums []int) []int {
	l := len(nums)
	arr := make([]int, l)
	copy(arr, nums)

	for i := 0; i < l-1; i++ {
		idx := i

		for j := i + 1; j < l; j++ {
			if comparator(arr[idx], arr[j]) {
				arr[idx], arr[j] = arr[j], arr[idx]
			}
		}
	}

	return arr
}
