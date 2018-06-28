package clusterer

func min_int(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max_int(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sqr(a float64) float64 {
	return a*a
}

func indexOf(elem int, slice []int) int {
	for i, e := range slice {
		if e == elem {
			return i
		}
	}
	return -1
}