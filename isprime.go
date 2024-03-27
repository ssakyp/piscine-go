/*

INSTRUCTION

Write a function that returns true if the int passed as parameter is a prime number.

Otherwise it returns false.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

(We also consider that 1 is not a prime number)

EXPECTED FUNCTION

package piscine

func IsPrime(nb, int) bool {

}


USAGE

package main

import "fmt"

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}

OUTPUT
true
false

*/

package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	// for loop is used to count starting from 2 up to nb itself and checks if nb % i if there is reminder that means there is a divider.

	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}

	return true
}
