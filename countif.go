package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, e := range tab {
		if f(e) == true {
			counter += 1
		}
	}
	return counter
}
