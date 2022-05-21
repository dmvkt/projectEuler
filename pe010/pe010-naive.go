/*
Project Euler, Problem 10:

Find the sum of all the primes below two million.

https://projecteuler.net/problem=10
*/

/*
Implementation note:
The code is almost the same as for PE007.
*/

package main

import (
	"fmt"
	"math"
)

const upper = 2000000

func main() {
	result := 2 + 3
	primes := []int{2, 3}
	number := primes[len(primes)-1]
	for number < upper {
		number += 2
		sqrt := int(math.Sqrt(float64(number)))
		isPrime := true
		for _, p := range primes {
			if p > sqrt {
				break
			}
			if number%p == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, number)
			result += number
		}
	}
	fmt.Println(result)
}
