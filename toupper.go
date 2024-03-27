package piscine

func ToUpper(s string) string {
	result := make([]rune, len(s))
	counter := 0

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result[counter] = char - 32
			counter++
		} else {
			result[counter] = char
			counter++
		}
	}
	return string(result[:counter])
}
