package piscine

func CollatzCountdown(start int) int {
	steps := 0

	if start <= 0 {
		return -1
	}

	for start != 1 {
		if isEven(start) == true {
			start /= 2
			steps++
		} else if isEven(start) == false {
			start = start*3 + 1
			steps++
		}
	}

	return steps
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}

	return false
}
