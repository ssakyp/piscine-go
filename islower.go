package piscine

func IsLower(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
		} else {
			return false
		}
	}
	return true
}
