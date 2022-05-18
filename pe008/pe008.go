/*
Project Euler, Problem 8:

Find the thirteen adjacent digits in the 1000-digit number (input.txt) that have
the greatest product. What is the value of this product?

https://projecteuler.net/problem=8
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputfile = "input.txt"
const nDigits = 13

var data string

func assertNoError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getData(filename string) {
	f, err := os.Open(filename)
	assertNoError(err)
	input := bufio.NewScanner(f)
	for input.Scan() {
		data += string(input.Text())
	}
	f.Close()
}

func main() {
	getData(inputfile)
	var max int
	cnt, prod := 0, 1
	for i, d := range data {
		digit, err := strconv.Atoi(string(d))
		assertNoError(err)
		if digit == 0 {
			cnt, prod = 0, 1
			continue
		}
		cnt++
		prod *= digit
		if cnt < nDigits {
			continue
		}
		if cnt > nDigits {
			prevDigit, _ := strconv.Atoi(string(data[i-nDigits]))
			prod /= prevDigit
		}
		if prod > max {
			max = prod
		}
	}
	fmt.Println(max)
}
