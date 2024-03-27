package piscine

func IterativeFactorial(x int) int {
	result := 1
	if x < 0 {
		return 0
	}
	if x == 0 {
		return 1
	}
	for i := 1; i <= x; i++ {
		result *= i
		if result < 0 {
			return 0
		}
	}
	return result
}
