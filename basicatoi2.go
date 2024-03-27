package piscine

func BasicAtoi2(s string) int {
	a := 0

	for _, char := range s {
		digit := int(char - '0')

		if digit >= 0 && digit <= 9 {
			a = a*10 + digit
		} else {
			return 0
		}
	}
	return a
}
