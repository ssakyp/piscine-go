package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	runeslice := []rune(s)
	for _, r := range runeslice {
		z01.PrintRune(r)
	}
}
