#!/usr/bin/python

"""
Project Euler, Problem 21:

Evaluate the sum of all the amicable numbers under 10000.

https://projecteuler.net/problem=21
"""

from math import sqrt

upper = 10000

def sum_divisors(n):
    result = 1
    for i in range(2, int(sqrt(n)) + 1):
        if n % i == 0:
            result += i
            if int(n / i) != i:
                result += int(n / i)
    return result

def main():
    divsums = [0]
    for i in range(1, upper + 1):
        divsums.append(sum_divisors(i))
    
    result = 0
    for i in range(1, len(divsums)):
        j = divsums[i]
        if j < i and divsums[j] == i:
            result += i + j
    print(result)

if __name__ == "__main__":
    main()
