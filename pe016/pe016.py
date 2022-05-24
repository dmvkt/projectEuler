#!/usr/bin/python

"""
Project Euler, Problem 16:

What is the sum of the digits of the number 2^1000?

https://projecteuler.net/problem=16
"""

def main():
    number = str(2**1000)
    result = 0
    for digit in number:
        result += int(digit)
    print(result)

if __name__ == "__main__":
    main()
