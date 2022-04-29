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
    fib_prev = 1
    fib_cur = 1
    
    while True:
        fib_new = fib_prev + fib_cur
        fib_prev, fib_cur = fib_cur, fib_new
        if fib_new > max_number:
            break
        if (fib_new % 2 == 0):
            result += fib_new
    return str(result)

if __name__ == "__main__":
    print(main())
