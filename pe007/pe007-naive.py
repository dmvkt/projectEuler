#!/usr/bin/python

"""
Project Euler, Problem 7:

What is the 10001st prime number?

https://projecteuler.net/problem=7
"""

from math import sqrt

count = 10001

primes = []

def main():
    primes = [3]
    number = primes[-1] + 2
    number_sqrt = sqrt(number)
    while len(primes) + 1 < count:
        not_prime = 0
        for p in primes:
            if p > number_sqrt:
                break
            if number % p == 0:
                not_prime = 1
                break
        if not_prime == 0:
            primes.append(number)
        number += 2
        number_sqrt = sqrt(number)
    return primes[-1]

if __name__ == "__main__":
    print(main())
