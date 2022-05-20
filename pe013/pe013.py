#!/usr/bin/python

"""
Project Euler, Problem 13:

Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.

https://projecteuler.net/problem=13
"""

import sys

data = []

def read_data(inputfile):
    global data
    try:
        with open(inputfile) as fi:
            for l in fi:
                data.append(l)
            fi.close()
    except:
        sys.stderr.write("Error: failed to read data from " + inputfile + "!")
        sys.exit(1)

def main():
    result = 0
    read_data("./input.txt")
    try:
        for d in data:
            result += int(d)
    except:
        sys.stderr.write("Error: invalid input data ('" + d + "')")
    return str(result)[0:10]

if __name__ == "__main__":
    print(main())
