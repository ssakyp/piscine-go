package piscine

func isAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func Capitalize(s string) string {
	words := []byte(s)
	capitalizeNext := true

	for i := 0; i < len(words); i++ {
		if isAlphanumeric(words[i]) {
			if capitalizeNext {
				if words[i] >= 'a' && words[i] <= 'z' {
					words[i] -= 'a' - 'A'
				}
				capitalizeNext = false
			} else if words[i] >= 'A' && words[i] <= 'Z' {
				words[i] += 'a' - 'A'
			}
		} else {
			capitalizeNext = true
		}
	}

	return string(words)
}
