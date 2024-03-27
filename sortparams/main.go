package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var args []string
	for _, s := range os.Args[1:] {
		args = append(args, s)
	}
	SortIt(args)
	for _, s := range args {
		for _, c := range s {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}

func SortIt(table []string) {
	for i := 0; i < len(table)-1; i++ {
		for j := 1; j < len(table); j++ {
			if table[j-1] > table[j] {
				temporary := table[j]
				table[j] = table[j-1]
				table[j-1] = temporary
			}
		}
	}
}
