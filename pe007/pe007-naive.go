/*
Project Euler, Problem 7:

What is the 10001st prime number?

https://projecteuler.net/problem=7
*/

package main

import (
	"fmt"
	"math"
)

const theNumber = 10001

func main() {
	primes := []int{2, 3}
	cnt := len(primes)
	number := primes[len(primes)-1]
	for cnt < theNumber {
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
			cnt++
		}
	}
	fmt.Println(number)
}
