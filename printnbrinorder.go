package piscine

import "github.com/01-edu/z01"

func intToDigits(n int) []int {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}

func Sort(arr []int) {
	count := len(arr)
	for i := count; i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				value := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = value
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := intToDigits(n)
	Sort(digits)

	for _, digit := range digits {
		z01.PrintRune(rune(digit) + '0')
	}
}
