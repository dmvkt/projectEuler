#!/usr/bin/python

"""
Project Euler, Problem 9:

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product a*b*c.

https://projecteuler.net/problem=9
"""

the_sum = 1000

def main():
    for a in range(1, the_sum):
        for b in range(a, the_sum):
            c = the_sum - a - b
            if c < b:
                break
            if a * a + b * b == c * c:
                return a * b *c

if __name__ == "__main__":
    print(main())

