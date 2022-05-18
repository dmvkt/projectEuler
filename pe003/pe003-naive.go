/*
Project Euler, Problem 3:

What is the largest prime factor of the number 600851475143?

https://projecteuler.net/problem=3
*/

package main

import "fmt"
import "math"

var number = 600851475143

func getPrimes(upper int) []int {
	var result []int
	var candidates = make([]uint8, upper+1)

	candidates[0], candidates[1] = 1, 1
	for i := 2; i <= upper; i++ {
		if candidates[i] == 1 {
			continue
		}
		for j := 2 * i; j <= upper; j += i {
			candidates[j] = 1
		}
	}

	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] == 1 {
			continue
		}
		result = append(result, i)
	}
	return result
}

func main() {
	result := 0
	var primes []int = getPrimes(int(math.Sqrt(float64(number))))
	for _, p := range primes {
		if number%p == 0 {
			result = p
			break
		}
	}
	fmt.Println(result)
}
