/*
Project Euler, Problem 2:

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.

https://projecteuler.net/problem=2
*/

package main

import "fmt"

var upper = 4 * 1000 * 1000

func main() {
	var sum int
	var fib2 int

	fib0, fib1 := 1, 1
	for {
		fib2 = fib0 + fib1
		fib0, fib1 = fib1, fib2
		if fib2 > upper {
			break
		}
		if fib2%2 == 0 {
			sum += fib2
		}
	}
	fmt.Println(sum)
}
