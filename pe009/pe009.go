/*
Project Euler, Problem 9:

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product a*b*c.

https://projecteuler.net/problem=9
*/

package main

import "fmt"

const sum = 1000

func main() {
	for a := 1; a < sum; a++ {
		for b := a; b < sum; b++ {
			c := sum - a - b
			if c < b {
				break
			}
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
