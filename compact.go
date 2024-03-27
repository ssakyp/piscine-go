package piscine

func Compact(ptr *[]string) int {
	count := 0
	var b []string

	for _, c := range *ptr {
		if c != "" {
			b = append(b, c)
			count++
		}
	}
	*ptr = b
	return count
}
