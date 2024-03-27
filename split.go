package piscine

func Split(s, sep string) []string {
	var strings []string
	var count int

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:len(sep)+i] == sep {
			strings = append(strings, s[i-count:i])
			count = -len(sep) + 1

		} else if i == len(s)-len(sep) {
			strings = append(strings, s[i-count:i+len(sep)])
		} else {
			count++
		}
	}
	return strings
}
