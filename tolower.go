package piscine

func ToLower(s string) string {
	result := make([]rune, len(s))
	counter := 0

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result[counter] = char + 32
			counter++
		} else {
			result[counter] = char
			counter++
		}
	}
	return string(result[:counter])
}
