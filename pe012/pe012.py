#!/usr/bin/python

"""
Project Euler, Problem 12:

What is the value of the first triangle number to have over five hundred divisors?

https://projecteuler.net/problem=12
"""

from math import sqrt

def count_divisors(n):
    cnt = 0
    for i in range(1, int(sqrt(n)) + 1):
        if n % i == 0:
            cnt += 1
            if int(n / i) != i:
                cnt += 1
    return cnt

def main():
    i, t = 1, 1
    max_cnt = 0
    while max_cnt <= 500:
        i += 1
        t += i
        cnt = count_divisors(t)
        if cnt > max_cnt:
            max_cnt = cnt
    print(t)

if __name__ == "__main__":
    main()
