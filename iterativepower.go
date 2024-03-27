/*
Write an iterative function that returns the value of nb to the power of power.
Negative numbers will return 0. Overflows do not have to be dealt with.

Expected function:

package piscine

func IterativePower(nb int, power int) int  {

}

Usage:
Here is a possible program to test your function:

package main

import "fmt"

func main() {
	fmt.Println(IterativePower(-2, 3))
}

And its output:

64

*/

package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	result := nb

	for i := 1; i < power; i++ {
		result *= nb
	}

	return result
}
