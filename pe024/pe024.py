#!/usr/bin/python

"""
Project Euler, Problem 24:

What is the millionth lexicographic permutation of the digits 0..9?

https://projecteuler.net/problem=24
"""

"""
Implementation note:
Each permutation can be described as:
 * an ordered arrangement (e.g. 9876543210),
 * an index, i.e. number of a permution in the lexicographic order,
 * in a Lehmer code (an order number in factorial number system).
"""

ndigits = 10
index = 1000000

def fact(n):
    if n == 1:
        return 1
    return n * fact(n - 1)

def index_to_lehmer(n):
    res = [0] * ndigits
    d = fact(ndigits - 1)
    n -= 1
    for i in range(1, ndigits):
        r = int(n / d)
        n -= r * d
        d /= ndigits - i
        res[i - 1] = r
    return res

def lehmer_to_arrangement(lehmer):
    res = []
    digits = [ str(i) for i in range(0, ndigits) ]
    for l in lehmer:
    	res.append(digits.pop(l))
    return res
    
def main():
    print(''.join(lehmer_to_arrangement(index_to_lehmer(index))))

if __name__ == "__main__":
    main()
