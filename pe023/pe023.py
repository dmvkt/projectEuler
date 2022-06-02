#!/usr/bin/python

"""
Project Euler, Problem 23:

A number n is called abundant if the sum of its proper divisors is greater 
than n. Given that all integers greater than 28123 can be written as the sum
of two abundant numbers, find the sum of all the positive integers which cannot
be written as the sum of two abundant numbers.

https://projecteuler.net/problem=23
"""

from math import sqrt

upper = 28123

def sum_divisors(n):
    result = 1
    for i in range(2, int(sqrt(n)) + 1):
        if n % i == 0:
            result += i
            if int(n / i) != i:
                result += int(n / i)
    return result

def main():
    abundant_numbers = []
    for i in range(1, upper + 1):
        if sum_divisors(i) > i:
            abundant_numbers.append(i)
    
    maxsum = 2 * abundant_numbers[-1]
    sums = [0] * (maxsum + 1)
    for i in abundant_numbers:
        for j in abundant_numbers:
            sums[i + j] = 1

    result = 0
    for i in range(1, upper + 1):
        if i > maxsum or sums[i] == 0:
            result += i
    print(result)

if __name__ == "__main__":
    main()
