#!/usr/bin/python

"""
Project Euler, Problem 4:

Find the largest palindrome made from the product of two 3-digit numbers.

https://projecteuler.net/problem=4
"""

def main():
    max_product = 0
    for x in range(100, 1000):
        for y in range(100, 1000):
            product = x * y
            if str(product) != str(product)[::-1]:
                continue
            if product > max_product:
                max_product = product
    return max_product

if __name__ == "__main__":
    print(main())
