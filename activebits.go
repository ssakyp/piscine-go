package piscine

func ActiveBits(n int) int {
	count := 0

	// Loop through each bit of the integer
	for n != 0 {
		// Check if the least significant bit (rightmost bit) is 1
		if n&1 == 1 {
			count++
		}

		// Right shift the integer by 1 bit
		n >>= 1
	}

	return count
}
