package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a[len(a)-1]
}
