package piscine

func FirstRune(s string) rune {
	r := []rune(s)
	for index := range r {
		return r[index]
	}
	return 0
}
