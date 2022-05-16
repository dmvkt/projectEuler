#!/usr/bin/python

"""
Project Euler, Problem 5:

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

https://projecteuler.net/problem=5
"""

from functools import reduce

max_number = 20

def get_gcd2(a, b):
    while b > 0:
        b, a = a % b, b
    return a

def get_gcd(numbers):
    return reduce(get_gcd2, numbers)

def get_lcm2(a, b):
    return (a * b) / get_gcd2(a, b)

def get_lcm(numbers):
    return reduce(get_lcm2, numbers)

def main():
    print(int(get_lcm(range(1, max_number + 1))))

if __name__ == "__main__":
    main()

