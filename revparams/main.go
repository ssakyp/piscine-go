package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Iterate over command-line arguments in reverse order
	for i := len(os.Args) - 1; i > 0; i-- {
		// Iterate over each character in the current argument
		for j := 0; j < len(os.Args[i]); j++ {
			char := rune(os.Args[i][j])
			// Print the character
			z01.PrintRune(char)
		}

		// Print a newline after each word
		z01.PrintRune('\n')
	}
}
