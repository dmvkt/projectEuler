#!/usr/bin/python

"""
Project Euler, Problem 26:

Find the value of d < 1000 for which 1/d contains the longest recurring
cycle (period) in its decimal fraction part.

https://projecteuler.net/problem=26
"""

"""
Implementation note:
1) Division/multiplication by 2, 5, 10 (divisors of 10) doesn't change the length
   of a period.
2) If d is not divisible by 2 and 5, then the period of 1/d is equal to p, where
   p is the smallest number such that 10 to the power of p has remainder 1 when
   divided by d: 10**p % d == 1.
3) The fractions of a form 1/p, where p is prime, have the longest periods.
"""

def get_period_len(d):
    # Remove powers of 2
    while d % 2 == 0:
        d /= 2
    # Remove powers of 5
    while d % 5 == 0:
        d /= 5
    if d == 1:
        return 0

    # Find the smallest power of 10 which gives remainder 1 when divided by d
    p = 1
    n = 10
    while True:
        n = n % d
        if n == 1:
            return p
        if n == d - 1:
            # n is equal to (-1) % d, thus (n*n) % d is 1
            return 2 * p
        p += 1
        n *= 10

def main():
    argmax, maxval = 0, 0
    # It is sufficient to consider only primes, but we'll check all odd numbers,
    # because it's simple and fast enough
    for i in range(3, 1000, 2):
        val = get_period_len(i)
        if val > maxval:
            maxval = val
            argmax = i
    print(argmax)

if __name__ == "__main__":
    main()
