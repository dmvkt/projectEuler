/*
Project Euler, Problem 20:

Find the sum of the digits in the number 100!

https://projecteuler.net/problem=20
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var sum int
	var bigint big.Int
	for _, s := range bigint.MulRange(1, 100).String() {
		d, _ := strconv.Atoi(string(s))
		sum += d
	}
	fmt.Println(sum)
}
