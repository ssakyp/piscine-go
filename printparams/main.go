package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, argument := range os.Args[1:] {
		// Iterate over each character in the current argument
		for _, char := range argument {
			if char == ' ' {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(char)
			}
		}
		// Print a newline after each argument
		z01.PrintRune('\n')
	}
}
