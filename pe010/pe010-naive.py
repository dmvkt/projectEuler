#!/usr/bin/python

"""
Project Euler, Problem 10:

Find the sum of all the primes below two million.

https://projecteuler.net/problem=10
"""

"""
Implementation note:
The code is almost the same as for PE007.
"""

from math import sqrt

upper = 2000000

def main():
    result = 2 + 3
    primes = [2, 3]
    number = primes[-1] + 2
    number_sqrt = sqrt(number)
    while number < upper:
        not_prime = 0
        for p in primes:
            if p > number_sqrt:
                break
            if number % p == 0:
                not_prime = 1
                break
        if not_prime == 0:
            primes.append(number)
            result += number
        number += 2
        number_sqrt = sqrt(number)
    return result

if __name__ == "__main__":
    print(main())
