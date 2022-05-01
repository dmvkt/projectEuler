#!/usr/bin/python

"""
Project Euler, Problem 4:

Find the largest palindrome made from the product of two 3-digit numbers.

https://projecteuler.net/problem=4
"""

"""
Implementation note:
Any palindrome of a form ABCCBA = A*100001 + B*10010 + C*1100 must be divisible
by GCD(100001, 10010, 1100) = 11.
Thus we know that the first number must be a multiple of 11.
"""

def main():
    max_product = 0
    divisor = 11
    for x in range(int((100 + divisor - 1)/divisor) * divisor, 1000, divisor):
        for y in range(100, 1000):
            product = x * y
            if str(product) != str(product)[::-1]:
                continue
            if product > max_product:
                max_product = product
    return max_product

if __name__ == "__main__":
    print(main())
