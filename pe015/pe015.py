#!/usr/bin/python

"""
Project Euler, Problem 15:

Starting in the top left corner of a 2×2 grid, and only being able to move to
the right and down, there are exactly 6 routes to the bottom right corner.
How many such routes are there through a 20×20 grid?

https://projecteuler.net/problem=15
"""

"""
Implementation note:
Each route has (in some specific order) 20 moves to the right and 20 moves
down, 40 moves in total. Thus, each route can be written as a sequence of
letters R and D: R for the move to the right, D for the move down.
The total number of such routes is equal to the number of words of length
40 with 20 R's and 20 D's. This is a well known binomial coefficient
that is equal to 40!/(20!*20!).
"""

gridX = 20
gridY = 20

def factorial(n):
    if n < 2:
        return 1
    return n * factorial(n - 1)

def main():
    print(int(factorial(gridX + gridY)/factorial(gridX)/factorial(gridY)))

if __name__ == "__main__":
    main()
