#!/usr/bin/python

"""
Project Euler, Problem 8:

Find the thirteen adjacent digits in the 1000-digit number (input.txt) that have
the greatest product. What is the value of this product?

https://projecteuler.net/problem=8
"""

import sys

mult_need = 13
data = ''

def read_data(inputfile):
    global data
    try:
        with open(inputfile, 'r') as fi:
            for l in fi:
                data += l[:-1]
            fi.close()
    except:
        sys.stderr.write("Error: input file '" + inputfile + "' not found\n")
        sys.exit(1)

def process_data():
    max_prod = 0
    prod = 1
    mult_count = 0
    pos = 0
    while pos < len(data):
        digit = int(data[pos])
        pos += 1
        if digit == 0:
            mult_count = 0
            prod = 1
            continue
        mult_count += 1
        prod *= digit
        if mult_count < mult_need:
            continue
        if mult_count > mult_need:
            prod = int(prod / int(data[pos - mult_need - 1]))
        if prod > max_prod:
            max_prod = prod
    return max_prod

def main():
    read_data('./input.txt')
    return process_data()

if __name__ == '__main__':
    print(main()
)
