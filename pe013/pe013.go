/*
Project Euler, Problem 13:

Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.

https://projecteuler.net/problem=13
*/

package main


import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var sum big.Int

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		var num big.Int
		num.SetString(string(input.Text()), 10)
		sum.Add(&sum, &num)

	}
	fmt.Println(sum.Text(10)[0:10])
}
