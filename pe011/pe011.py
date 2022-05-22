#!/usr/bin/python

"""
Project Euler, Problem 11:

What is the greatest product of four adjacent numbers in the same direction
(up, down, left, right, or diagonally) in the 20×20 grid (input.txt)?

https://projecteuler.net/problem=11
"""

"""
Implementation note:
The solution could have been more compact and straightforward, but the goal was
to use iterators.
"""

import sys

class IterColumns():
    """
    Iterate over columns in a 2D array; insert 0 on the edge.
    """
    def __init__(self, data):
        self.data = data

    def __iter__(self):
        self.i = -1
        self.j = 0
        return self
    
    def __next__(self):
        self.i += 1
        if self.i == len(self.data):
            return 0
        if self.i > len(self.data):
            self.i = 0
            self.j += 1
        if self.j == len(self.data[0]):
            raise StopIteration
        return self.data[self.i][self.j]

class IterRows():
    """
    Iterate over rows in a 2D array; insert 0 on the edge.
    """
    def __init__(self, data):
        self.data = data

    def __iter__(self):
        self.i = 0
        self.j = -1
        return self
    
    def __next__(self):
        self.j += 1
        if self.j == len(self.data[0]):
            return 0
        if self.j > len(self.data[0]):
            self.i += 1
            self.j = 0
        if self.i == len(self.data):
            raise StopIteration
        return self.data[self.i][self.j]

class IterDiagonals():
    """
    Iterate over primary diagonals (\) in a 2D array; insert 0 on the edge.
    """
    def __init__(self, data):
        self.data = data

    def __iter__(self):
        self.k = -(len(self.data))
        self.new_k_next = True
        return self
    
    def shift_k(self):
        self.k += 1
        if self.k >= len(self.data[0]):
            raise StopIteration
        self.new_k_next = False
        if self.k < 0:
            self.i = -self.k - 1
            self.j = -1
        else:
            self.i = -1
            self.j = self.k - 1
    
    def __next__(self):
        if self.new_k_next:
            self.shift_k()
        self.i += 1
        self.j += 1
        if self.i >= len(data) or self.j >= len(data[0]):
            self.new_k_next = True
            return 0
        return self.data[self.i][self.j]

class IterSDiagonals():
    """
    Iterate over secondary diagonals (/) in a 2D array; insert 0 on the edge.
    """
    def __init__(self, data):
        self.data = data

    def __iter__(self):
        self.k = -(len(self.data))
        self.new_k_next = True
        return self
    
    def shift_k(self):
        self.k += 1
        if self.k >= len(self.data[0]):
            raise StopIteration
        self.new_k_next = False
        if self.k < 0:
            self.i = -self.k - 1
            self.j = len(self.data[0])
        else:
            self.i = -1 
            self.j = len(self.data[0]) - self.k
    
    def __next__(self):
        if self.new_k_next:
            self.shift_k()
        self.i += 1
        self.j -= 1
        if self.i >= len(self.data) or self.j < 0:
            self.new_k_next = True
            return 0
        return self.data[self.i][self.j]

data = []

def read_data(filename):
    try:
        with open(filename, 'r') as fi:
            for l in fi:
                if l.strip() == "":
                    continue
                data.append([ int(x) for x in l.split(' ') ])
            fi.close()
    except IOError:
        sys.stderr.write("Error: failed to read data!\n")
        sys.exit(1)

def get_max_product(iterdata, n):
    """
    Find the greatest product of n adjacent numbers in a 1D array (passed as
    an iterator).
    """
    max_product = 0
    have, product = [], 1
    for value in iterdata:
        if value == 0:
            have, product = [], 1
            continue
        product *= value
        have.append(value)
        if len(have) < n:          
            continue
        if len(have) > n:
            product = int(product / have[-(n + 1)])
        if product > max_product:
            max_product = product
    return max_product

def main():
    read_data("input.txt")
    result = []
    for iterator in [IterRows, IterColumns, IterDiagonals, IterSDiagonals]:
        result.append(get_max_product(iterator(data), 4))
    print(max(result))

if __name__ == "__main__":
    main()
