package piscine

func Unmatch(a []int) int {
	// Create a map to store the frequency of each element
	freq := make(map[int]int)

	// Count the frequency of each element
	for _, num := range a {
		freq[num]++
	}

	// Iterate through the slice again to find the first element with odd frequency
	for _, num := range a {
		if freq[num]%2 != 0 {
			return num
		}
	}

	// If no element with odd frequency is found, return -1
	return -1
}
