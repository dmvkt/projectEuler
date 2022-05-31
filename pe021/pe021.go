/*
Project Euler, Problem 21:

Evaluate the sum of all the amicable numbers under 10000.

https://projecteuler.net/problem=21
*/

package main

import (
	"fmt"
	"math"
)

const upper = 10000

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
	divsums := []int{0}
	for i := 1; i <= upper; i++ {
		divsums = append(divsums, sumDivisors(i))
	}

	result := 0
	for i, d := range divsums {
		if d < i && divsums[d] == i {
			result += i + d
		}
	}
	fmt.Println(result)
}
