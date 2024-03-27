package piscine

func Map(f func(int) bool, a []int) []bool {
	var isPrime []bool
	for _, e := range a {
		isPrime = append(isPrime, f(e))
	}
	return isPrime
}
