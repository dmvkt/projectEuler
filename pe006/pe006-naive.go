/*
Project Euler, Problem 6:

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

https://projecteuler.net/problem=6
*/

package main

import "fmt"

const max = 100

func main() {
	sum, squares := 0, 0
	for i := 1; i <= max; i++ {
		sum += i
		squares += i * i
	}
	fmt.Println(sum*sum - squares)
}

