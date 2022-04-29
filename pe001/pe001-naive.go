/*
Project Euler, Problem 1:

Find the sum of all the multiples of 3 or 5 below 1000.

https://projecteuler.net/problem=1
*/

package main

import "fmt"

var divisors = []int{3, 5}
var max_number = 999

func main() {
    result := 0
    for i := 2; i <= max_number; i++ {
        for _, d := range divisors {
            if (i % d == 0) {
                result += i
                break
            }
        }
    }
    fmt.Println(result)
}

