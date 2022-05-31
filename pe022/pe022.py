#!/usr/bin/python

"""
Project Euler, Problem 22:

Begin by sorting the names in input.txt into alphabetical order.
Then working out the alphabetical value for each name, multiply this value
by its alphabetical position in the list to obtain a name score.
What is the total of all the name scores in the file?

https://projecteuler.net/problem=22
"""

def main():
    with open("input.txt", "r") as fi:
        for l in fi:
            data = [ x.strip('"') for x in l.split(',') ]
        fi.close()
    data.sort()

    total = 0
    i = 1
    for d in data:
        score = 0
        for char in d:
            score += ord(char) - ord('A') + 1
        total += i * score
        i += 1
    print(total)

if __name__ == "__main__":
    main()

