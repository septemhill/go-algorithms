package sorting

func BubbleSort(nums []int) []int {
	l := len(nums)
	arr := make([]int, l)
	copy(arr, nums)

	for i := 0; i < l-1; i++ {
		for j := 1; j < l; j++ {
			if comparator(arr[j-1], arr[j]) {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	return arr
}
