/*
Project Euler, Problem 5:

What is the smallest positive number that is evenly divisible by all of
the numbers from 1 to 20?

https://projecteuler.net/problem=5
*/

package main

import "fmt"

const max = 20

type reduceFunc func(int, int) int

func reduce(f reduceFunc, numbers []int) int {
	result := numbers[0]
	for _, n := range numbers {
		result = f(result, n)
	}
	return result
}

func getGCD2(a, b int) int {
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func getLCM2(a, b int) int {
	return int(a * b / getGCD2(a, b))
}

func getGCD(numbers []int) int {
	return reduce(getGCD2, numbers)
}

func getLCM(numbers []int) int {
	return reduce(getLCM2, numbers)
}

func main() {
	var numbers []int
	for i := 2; i <= max; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(getLCM(numbers))
}
