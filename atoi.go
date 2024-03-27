package piscine

func Atoi(s string) int {
	final := 0
	multiplier := 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' && i == 0 {
			return final * -1
		} else if s[i] == '+' && i == 0 {
			return final
		}
		if !(s[i] >= '0' && s[i] <= '9') {
			return 0
		}

		final += multiplier * int(s[i]-'0')

		multiplier *= 10
	}
	return final
}
