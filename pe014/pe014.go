/*
Project Euler, Problem 14:

For the Collatz sequence, which starting number, under one million, produces
the longest chain?

https://projecteuler.net/problem=14
*/

package main

import (
	"fmt"
)

const upper = 1000000

var cache [upper]int

func getNext(n int) int {
	if n&1 == 0 {
		return n >> 1
	}
	return 3*n + 1
}

func getChainLen(n int) int {
	i := n
	chainLen := 0
	chain := []int{i}
	for i > 1 {
		i = getNext(i)
		chain = append(chain, i)
		if i < upper && cache[i] != 0 {
			chainLen = len(chain) + cache[i]
			break
		}
	}
	if chainLen == 0 {
		chainLen = len(chain)
	}
	for i, c := range chain {
		if c >= upper {
			continue
		}
		if cache[c] > 0 {
			break
		}
		cache[c] = chainLen - i
	}
	return chainLen
}

func main() {
	maxLen := 0
	maxVal := 0
	for n := 2; n < upper; n++ {
		chainLen := getChainLen(n)
		if chainLen > maxLen {
			maxLen = chainLen
			maxVal = n
		}
	}
	fmt.Println(maxVal)
}
