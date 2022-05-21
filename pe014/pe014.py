#!/usr/bin/python

"""
Project Euler, Problem 14:

For the Collatz sequence, which starting number, under one million, produces
the longest chain?

https://projecteuler.net/problem=14
"""

upper = 1000000
cache = [0] * upper

def get_next(n):
    if n % 2 == 0:
        return int(n / 2)
    else:
        return 3 * n + 1

def get_chainlen(n):
    i = n
    chainlen = 0
    chain = [i]
    while i != 1:
        i = get_next(i)
        chain.append(i)
        if i < upper and cache[i] != 0:
            chainlen = cache[i] + len(chain)
            break
    if chainlen == 0:
        chainlen = len(chain)
    dist = -1
    for c in chain:
        dist += 1
        if c >= upper:
            continue
        if cache[c] != 0:
            break
        cache[c] = chainlen - dist
    return chainlen

def main():
    max_val = 0
    max_len = 0
    n = 2
    while n < upper:
        chainlen = get_chainlen(n)
        if chainlen > max_len:
            max_len = chainlen
            max_val = n
        n += 1
    print(max_val)

if __name__ == "__main__":
    main()
