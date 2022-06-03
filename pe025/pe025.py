#!/usr/bin/python

"""
Project Euler, Problem 25:

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

https://projecteuler.net/problem=25
"""

"""
Implementation note:

There is a formula to calculate the nth term in the Fibonacci sequence:
Fib(n) = (P**n - Q**n)/sqrt(5),
P = (1 + sqrt(5))/2, Q = (1 - sqrt(5))/2
Fib(0) = 0, Fib(1) = 1

Since Q < 1, the term Q**n/sqrt(5) can be ignored for sufficiently large n.
In fact, for n > 35 we can use the formula:
Fib2(n) = int(P**n / sqrt(5))

Thus, the number of digits in the nth Fibonacci term is given by:
1 + int(log10(Fib2(n))) = 1 + int(log10(int(P**n / sqrt(5)))) = 1 + int(n*log10(P) - log10(5)/2)

Now we can get the inequality for the terms containing at least D digits:
1 + n*log10(P) - log10(5)/2 >= D
=> n >= (D - 1 + log10(5)/2) / log10(P)
"""

from math import sqrt, log10, ceil

P = (1 + sqrt(5))/2

def fib_index(D):
    """
    Return the index of the first Fibonacci term to contain D digits
    """
    if D <= 1:
        return D
    return ceil((D - 1 + log10(5)/2) / log10(P))
    
if __name__ == "__main__":
    print(fib_index(1000))

