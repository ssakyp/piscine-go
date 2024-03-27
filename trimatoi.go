package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	started := false

	for _, char := range s {
		if char == '-' && !started {
			sign = -1
		} else if char == '+' && !started {
			sign = 1
		} else if !(char >= '0' && char <= '9') {
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
			started = true
		}
	}

	return result * sign
}
