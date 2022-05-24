/*
Project Euler, Problem 16:

What is the sum of the digits of the number 2^1000?

https://projecteuler.net/problem=16
*/

package main

import (
	"fmt"
	"math/big"
)

const power = 1000
const inBase = 2
const outBase = 10

func sumDigits(n *big.Int, base int) int {
	sum := 0
	for _, d := range n.Text(base) {
		sum += (int(d) - int('0'))
	}
	return sum
}

func main() {
	var n big.Int
	n.Exp(big.NewInt(inBase), big.NewInt(power), nil)
	fmt.Println(sumDigits(&n, outBase))
}
