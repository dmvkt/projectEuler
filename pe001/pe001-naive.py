#!/usr/bin/python

"""
Project Euler, Problem 1:

Find the sum of all the multiples of 3 or 5 below 1000.

https://projecteuler.net/problem=1
"""

divisors = [3, 5]
max_number = 999

def main():
    result = 0
    for i in range(2, max_number + 1):
        if min([ i % d for d in divisors ]) == 0:
            result += i   
    return str(result)

if __name__ == "__main__":
    print(main())
