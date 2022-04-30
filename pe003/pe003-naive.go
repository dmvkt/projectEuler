/*
Project Euler, Problem 3:

What is the largest prime factor of the number 600851475143?

https://projecteuler.net/problem=3
*/

package main

import "fmt"
import "math"

var the_number = 600851475143

func generate_primes(upper_bound int) []int {
    var candidates = make([]uint8, upper_bound + 1)
    candidates[0] = 1
    candidates[1] = 1
    for i := 2; i <= upper_bound; i++ {
        if candidates[i] == 1 {
            continue
        }
        for j := 2 * i; j <= upper_bound; j += i {
            candidates[j] = 1
        }
    }
    
    var result []int
    
    for i := len(candidates) - 1; i >= 0; i-- {
        if candidates[i] == 1 {
            continue
        }
        result = append(result, i)
    }
    return result
}

func main() {
    result := 0
    var primes []int = generate_primes(int(math.Sqrt(float64(the_number))))
    for _, p := range primes {
        if the_number % p == 0 {
            result = p
            break
        }
    }
	
	fmt.Println(result)
}
