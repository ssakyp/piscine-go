package piscine

func Compare(a, b string) int {
	a_byte := []byte(a)
	b_byte := []byte(b)
	if a == b {
		return 0
	}
	len_a := 0
	len_b := 0

	for range a_byte {
		len_a++
	}

	for range b_byte {
		len_b++
	}

	for i := 0; i < len_a && i < len_b; i++ {
		if a_byte[i] < b_byte[i] {
			return -1
		} else {
			return 1
		}
	}
	return 0
}
