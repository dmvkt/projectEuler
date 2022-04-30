#!/usr/bin/python

"""
Project Euler, Problem 3:

What is the largest prime factor of the number 600851475143?

https://projecteuler.net/problem=3
"""

from math import sqrt

the_number = 600851475143

primes = []

def generate_primes(upper_bound):
    candidates = [1] * (upper_bound + 1)
    candidates[0] = 0
    candidates[1] = 0
    for i in range(2, upper_bound + 1):
        if candidates[i] == 0:
            continue
        for j in range(2 * i, upper_bound + 1, i):
            candidates[j] = 0
    return [ i for i in range(len(candidates) - 1, -1, -1) if candidates[i] > 0 ]

def main():
    primes = generate_primes(int(sqrt(the_number)))
    for p in primes:
        if the_number % p == 0:
            return p

if __name__ == "__main__":
    print(main())
