#!/usr/bin/python

"""
Project Euler, Problem 20:

Find the sum of the digits in the number 100!

https://projecteuler.net/problem=20
"""

def fact(n):
    if n == 1:
        return 1
    return n * fact(n - 1)

print(sum([ int(s) for s in str(fact(100)) ]))
