/*
Project Euler, Problems 18 & 67:

Find the maximum total from top to bottom in a text file containing a triangle
with:
 1) 15 (input18.txt) rows,
 2) 100 (input67.txt) rows.

https://projecteuler.net/problem=18
https://projecteuler.net/problem=67
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func assertNoError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(filename string) {
	var totals []int
	f, err := os.Open(filename)
	assertNoError(err)
	input := bufio.NewScanner(f)

	k := 0
	for input.Scan() {
		var values []int
		for _, s := range strings.Split(input.Text(), " ") {
			value, err := strconv.Atoi(s)
			assertNoError(err)
			values = append(values, value)
		}

		k += 1
		if k == 1 {
			totals = append(totals, values[0])
		} else {
			totals = append(totals, values[k-1]+totals[k-2])
			for i := k - 2; i > 0; i-- {
				totals[i] = values[i] + max(totals[i], totals[i-1])
			}
			totals[0] += values[0]
		}
	}

	maxTotal := 0
	for _, x := range totals {
		if x > maxTotal {
			maxTotal = x
		}
	}
	fmt.Println(maxTotal)
}

func main() {
	run("input18.txt")
	run("input67.txt")
}
