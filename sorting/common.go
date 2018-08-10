package sorting

func intSliceConcat(slice ...[]int) []int {
	v := []int{}

	for _, e := range slice {
		v = append(v, e...)
	}

	return v
}

func comparator(i, j int) bool {
	return i > j
}
