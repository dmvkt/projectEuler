#!/usr/bin/python

"""
Project Euler, Problem 6:

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

https://projecteuler.net/problem=6
"""

max_number = 100

def main():
    the_sum = sum(list(range(1, max_number + 1)))
    squares = sum([ i * i for i in range(1, max_number + 1) ])
    print(the_sum * the_sum - squares)

if __name__ == "__main__":
    main()
