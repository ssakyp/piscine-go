package piscine

func RecursiveFactorial(x int) int {
	if x < 0 || x > 20 {
		return 0
	}
	if x == 0 {
		return 1
	}

	result := x * RecursiveFactorial(x-1)
	return result
}
