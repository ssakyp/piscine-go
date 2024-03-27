package piscine

func IsUpper(str string) bool {
	for index := range str {
		if str[index] >= 'A' && str[index] <= 'Z' {
		} else {
			return false
		}
	}
	return true
}
