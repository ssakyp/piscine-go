package piscine

func LastRune(s string) rune {
	r := []rune(s)

	counter := 0

	for range r {
		counter++
	}
	return r[counter-1]
}
