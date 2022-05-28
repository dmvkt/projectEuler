/*
Project Euler, Problem 28:

Starting with the number 1 and moving to the right in a clockwise direction
a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed
in the same way?

https://projecteuler.net/problem=28
*/

/*
Implementation note:
1) The spiral consists of square layers with side lengths 3, 5, 7 and so on.
   For example, the layer with side length 3:
     7  8  9
     6     2
     5  4  3
2) In the layer with side length k the top rigth element is k*k.
   The other corner elements are: k*k - (k-1), k*k - 2(k-1), k*k - 3(k-1).
   The sum of corner elements is J[k] = 4*k*k - (k-1)(1+2+3) = 4*k*k - 6(k-1).
3) The sum of the numbers on the diagonals in an N by N spiral is equal to
   the sum of corner elements in all layers up to N:
     S[N] = 1 + J[3] + J[5] + ... + J[N]
*/

package main

import "fmt"

func main() {
	sum := 1
	for k := 3; k <= 1001; k += 2 {
		sum += 4*k*k - 6*(k-1)
	}
	fmt.Println(sum)
}
