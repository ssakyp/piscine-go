package piscine

func Abort(a, b, c, d, e int) int {
	num := []int{a, b, c, d, e}

	// bubble sort
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				// Swap elements
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	result := num[2]
	return result
}
