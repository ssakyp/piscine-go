package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Initialize a boolean variable isName to false
	isName := false

	// Iterate over each character in the program's name (os.Args[0])
	for _, char := range os.Args[0] {
		// Check if the current character is '/'
		if char == '/' {
			// Set isName to true, indicating that the program name is expected
			isName = true
			continue // Skip to the next iteration
		}
		// Check if isName is true (indicating that we are in the program name part)
		if isName {
			// Print the current character using z01 package (character-wise printing)
			z01.PrintRune(char)
		}
	}
	// Print a newline character to separate the output from the shell prompt
	z01.PrintRune('\n')
}
