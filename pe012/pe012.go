/*
Project Euler, Problem 12:

What is the value of the first triangle number to have over five hundred divisors?

https://projecteuler.net/problem=12
*/

package main

import (
	"fmt"
	"math"
)

func countDivisors(n int) int {
	cnt := 0
	sqrt := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			cnt += 1
			if int(n/i) != i {
				cnt += 1
			}
		}
	}
	return cnt
}

func main() {
	i, t := 1, 1
	maxCnt := 0
	for maxCnt <= 500 {
		i += 1
		t += i
		cnt := countDivisors(t)
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}
	fmt.Println(t)
}
