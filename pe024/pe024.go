/*
Project Euler, Problem 24:

What is the millionth lexicographic permutation of the digits 0..9?

https://projecteuler.net/problem=24
*/

/*
Implementation note:
Each permutation can be described as:
 * an ordered arrangement (e.g. 9876543210),
 * an index, i.e. number of a permution in the lexicographic order,
 * in a Lehmer code (an order number in factorial number system).
*/

package main

import "fmt"

const ndigits = 10
const index = 1000000

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func indexToLehmer(n int) []int {
	res := make([]int, ndigits)
	d := fact(ndigits - 1)
	n -= 1
	for i := 1; i < ndigits; i++ {
		r := int(n / d)
		n -= r * d
		d /= ndigits - i
		res[i-1] = r
	}
	return res
}

func lehmerToArrangement(lehmer []int) string {
	var res []rune
	var digits []rune
	for i := 0; i < ndigits; i++ {
		digits = append(digits, rune(int('0')+i))
	}

	for _, l := range lehmer {
		res = append(res, digits[l])
		digits = append(digits[:l], digits[l+1:]...)
	}
	return string(res)
}

func main() {
	fmt.Println(lehmerToArrangement(indexToLehmer(index)))
}
