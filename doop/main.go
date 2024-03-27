package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	num1, ok1 := parseToInt(os.Args[1])
	operator := os.Args[2]
	num2, ok2 := parseToInt(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	switch operator {
	case "+":
		if (num1 > 0 && num2 > 0 && num1 > (1<<63-1)-num2) || (num1 < 0 && num2 < 0 && num1 < (-1<<63)-num2) {
			return
		}
		printInt(num1 + num2)
	case "-":
		if (num1 < 0 && num2 > 0 && num1 < (-1<<63)+num2) || (num1 > 0 && num2 < 0 && num1 > (1<<63-1)+num2) {
			return
		}
		printInt(num1 - num2)
	case "*":
		if num1 != 0 && num2 != 0 && (num1 > (1<<63-1)/num2 || num1 < (-1<<63)/num2) {
			return
		}
		printInt(num1 * num2)
	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		printInt(num1 / num2)
	case "%":
		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		printInt(num1 % num2)
	default:
		return
	}
}

func parseToInt(s string) (int64, bool) {
	var n int64
	var sign int64 = 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int64(s[i]-'0')
	}
	return sign * n, true
}

func printInt(n int64) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}

	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	}

	digits := []byte{}
	for n > 0 {
		digit := byte(n % 10)
		digits = append([]byte{digit + '0'}, digits...)
		n /= 10
	}

	os.Stdout.Write(digits)
	os.Stdout.WriteString("\n")
}
