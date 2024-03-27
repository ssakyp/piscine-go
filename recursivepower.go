/*
INSTRUCTIONS

Write a recursive function that returns the value of nb to the power of power.

Negative powers wil return 0.

Overflows do not have to be dealt with.

for is forbidden for this exercise.

EXPECTED FUNCTION:

package piscine

func RecursivePower(nb int, power int) int {

}



USAGE:

package main

import "fmt"

func main() {
	fmt.Println(RecursivePower(4, 3))
}




OUTPUT:

64

*/

package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}

	if power < 0 {
		return 0
	}

	result := nb * RecursivePower(nb, (power-1))

	return result
}
