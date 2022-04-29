/*
Project Euler, Problem 2:

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.

https://projecteuler.net/problem=2
*/

package main

import "fmt"

var max_number = 4 * 1000 * 1000

func main() {
	result := 0

	fib_prev := 1
	fib_cur := 1
	fib_new := 1

	for {
		fib_new = fib_prev + fib_cur
		fib_prev, fib_cur = fib_cur, fib_new
		if fib_new > max_number {
			break
		}
		if fib_new % 2 == 0 {
			result += fib_new
		}
	}
	fmt.Println(result)
}
