/*
INSTRUCTIONS

Write a recursive finction that returns the value at the position index

in the fibonacci sequence.

The first value is at the index 0.

The sequence starts this way: 0, 1, 1, 2, 3, etc...

A negative index will return -1.

for is forbidden for this exercise.

USAGE
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
}

OUTPUT

3

*/

package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}

	if index == 0 {
		return 0
	}

	if index == 1 {
		return 1
	}

	return Fibonacci(index-1) + Fibonacci(index-2)
}
