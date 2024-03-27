package piscine

func Any(f func(string) bool, a []string) bool {
	for _, e := range a {
		if f(e) == true {
			return true
		}
	}
	return false
}
