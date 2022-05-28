#!/usr/bin/python

"""
Project Euler, Problem 2:

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.

https://projecteuler.net/problem=2
"""

max_number = 4 * 10**6

def main():
    result = 0
    fib0 = 0
    fib1 = 1
    
    while True:
        fib2 = fib0 + fib1
        fib0, fib1 = fib1, fib2
        if fib2 > max_number:
            break
        if (fib2 % 2 == 0):
            result += fib2
    return str(result)

if __name__ == "__main__":
    print(main())
