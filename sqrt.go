/*
INSTRUCTIONS:

Write a function that returns the square root of the int passed as parameter, if that square root is a whole number.

Otherwise it returns 0.


EXPECTED FUNCTION:

package piscine

func Sqrt(nb, int) int {

}


USAGE:
package main

import "fmt"

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

OUTPUT:

2
0
*/

package piscine

func Sqrt(nb int) int {
	for i := 1; i <= nb; i++ {
		if nb/i == i && nb%i == 0 {
			return i
		}
	}
	return 0
}
