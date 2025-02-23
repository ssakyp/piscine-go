/*
Write a function that returns the first prime number that is equal or superior to the int passed as parameter.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

EXPECTED FUNCTION:

func FindNextPrime(nb int) int {

}

USAGE:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(5))
	fmt.Println(piscine.FindNextPrime(4))
}


OUTPUT:

$ go run .
5
5
$

*/

package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	for {
		isPrime := true
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			return nb
		}

		nb++
	}
}
