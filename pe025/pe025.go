/*
Project Euler, Problem 25:

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

https://projecteuler.net/problem=25
*/

/*
Implementation note:

There is a formula to calculate the nth term in the Fibonacci sequence:
Fib(n) = (P**n - Q**n)/sqrt(5),
P = (1 + sqrt(5))/2, Q = (1 - sqrt(5))/2
Fib(0) = 0, Fib(1) = 1

Since Q < 1, the term Q**n/sqrt(5) can be ignored for sufficiently large n.
In fact, for n > 35 we can use the formula:
Fib2(n) = int(P**n / sqrt(5))

Thus, the number of digits in the nth Fibonacci term is given by:
1 + int(log10(Fib2(n))) = 1 + int(log10(int(P**n / sqrt(5)))) = 1 + int(n*log10(P) - log10(5)/2)

Now we can get the inequality for the terms containing at least D digits:
1 + n*log10(P) - log10(5)/2 >= D
=> n >= (D - 1 + log10(5)/2) / log10(P)
*/

package main

import (
	"fmt"
	"math"
)

// Return the index of the first Fibonacci term to contain d digits
func fibIndex(d int) int {
	if d <= 1 {
		return d
	}
	index := (float64(d) - 1 + math.Log10(5)/2) / math.Log10((1+math.Sqrt(5))/2)
	return int(math.Ceil(index))
}

func main() {
	fmt.Println(fibIndex(1000))
}
