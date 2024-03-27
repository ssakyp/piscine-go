package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascendingOrder := true
	descendingOrder := true

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			ascendingOrder = false
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			descendingOrder = false
		}
	}

	return ascendingOrder || descendingOrder
}

func f(a, b int) int {
	if a < b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
