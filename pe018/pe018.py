#!/usr/bin/python

"""
Project Euler, Problems 18 & 67:

Find the maximum total from top to bottom in a text file containing a triangle
with:
 1) 15 (input18.txt) rows,
 2) 100 (input67.txt) rows.

https://projecteuler.net/problem=18
https://projecteuler.net/problem=67
"""

"""
Implementation note:
Denote x[k,i] as the i-th element in the k-th row, and total[k,i] as
the partial total from top to the i-th element in the k-th row.
The total[] can be evaluated as follows:
 * total[k, 0] = x[k, 0] + total[k-1, 0]
 * total[k, k] = x[k, k] + total[k-1, k-1]
 * total[k, i] = x[k, i] + max(total[k-1, i-1], total[k-1, i])

Note that to find total[k, *] we only need total[k-1, *], so there
is no need to keep all previous rows. Thus, memory footprint is
proportional only to the maximal length of a row.
"""

import sys

def main(filename):
    totals = []
    k = 0
    try:
        with open(filename, "r") as fi:
            for l in fi:
                k += 1
                if k == 1:
                    totals.append(0)
                else:
                    totals.append(totals[k - 2])
                    for i in range(k - 2, 0, -1):
                        totals[i] = max(totals[i], totals[i - 1])

                values = [ int(x) for x in l.split(' ') ]
                for i in range(0, k):
                    totals[i] += values[i]
            fi.close()
        print(max(totals))
    except:
        sys.stderr.write("Error: failed to read data!\n")
        sys.exit(1)

if __name__ == "__main__":
    main("input18.txt")
    main("input67.txt")

