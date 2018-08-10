package sorting

func InsertionSort(nums []int) []int {
	l := len(nums)
	arr := make([]int, l)

	copy(arr, nums)

	return arr
}
