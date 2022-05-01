/*
Project Euler, Problem 4:

Find the largest palindrome made from the product of two 3-digit numbers.

https://projecteuler.net/problem=4
*/

/*
Implementation note:
Any palindrome of a form ABCCBA = A*100001 + B*10010 + C*1100 must be divisible
by GCD(100001, 10010, 1100) = 11.
Thus we know that the first number must be a multiple of 11.
*/

package main

import "fmt"
import "strconv"

func reverse_string(s string) string {
    r := []rune(s)
    for i, j := 0, len(s) - 1; i < int(len(s)/2); i, j = i + 1, j - 1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func main() {
    const divisor = 11
    max_product := 0
    for x := int((100 + divisor - 1)/divisor) * divisor; x < 1000; x += divisor {
        for y := 100; y < 1000; y++ {
            product := x * y
            if strconv.Itoa(product) != reverse_string(strconv.Itoa(product)) {
                continue
            }
            if product > max_product {
                max_product = product
            }
        }
    }
    fmt.Println(max_product)
}



