/*
Project Euler, Problem 23:

A number n is called abundant if the sum of its proper divisors is greater
than n. Given that all integers greater than 28123 can be written as the sum
of two abundant numbers, find the sum of all the positive integers which cannot
be written as the sum of two abundant numbers.

https://projecteuler.net/problem=23
*/

package main

import (
	"fmt"
	"math"
)

const upper = 28123

func sumDivisors(n int) int {
	sum := 1
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			sum += i
			if int(n/i) != i {
				sum += int(n / i)
			}
		}
	}
	return sum
}

func main() {
	var abunduntNumbers []int
	for i := 1; i <= upper; i++ {
		if sumDivisors(i) > i {
			abunduntNumbers = append(abunduntNumbers, i)
		}
	}

	maxsum := 2 * abunduntNumbers[len(abunduntNumbers)-1]

	sums := make([]bool, maxsum+1)
	for _, i := range abunduntNumbers {
		for _, j := range abunduntNumbers {
			sums[i+j] = true
		}
	}

	result := 0
	for i := 1; i <= upper; i++ {
		if i > maxsum || !sums[i] {
			result += i
		}
	}

	fmt.Println(result)
}
