#!/usr/bin/python

"""
Project Euler, Problem 17:

If all the numbers from 1 to 1000 (one thousand) inclusive were written out
in words, how many letters would be used? 

https://projecteuler.net/problem=17

"""

hyphen = "-"
space = " "

numbers = { 1 : "one", 2 : "two", 3 : "three", 4 : "four", 5 : "five", 6 : "six",
    7 : "seven", 8: "eight", 9 : "nine", 10 : "ten", 11 : "eleven", 12 : "twelve",
    13 : "thirteen", 14: "fourteen", 15 : "fifteen", 16 : "sixteen",
    17 : "seventeen", 18 : "eighteen", 19 : "nineteen", 20 : "twenty",
    30 : "thirty", 40 : "forty", 50 : "fifty", 60 : "sixty", 70 : "seventy",
    80 : "eighty", 90 : "ninety" }

def number_to_text(n):
    global space
    global hyphen
    result = ""
    
    if n >= 1000:
        x = int( (n % 10000) / 1000 )
        n -= x * 1000
        result += numbers[x] + space + "thousand"
    if n >= 100:
        if len(result) > 0:
            result += space
        x = int( (n % 1000) / 100 )
        n -= x * 100
        result += numbers[x] + space + "hundred"
    if n > 0:
        if len(result) > 0:
            result += space + "and" + space
        if n < 20 or n % 10 == 0:
            result += numbers[n]
        else:
            result += numbers[int(n / 10)  * 10] + hyphen + numbers[n % 10]
    return result

if __name__ == "__main__":
    space = ""
    hyphen = ""
    
    result = 0
    for i in range(1, 1001):
        s = number_to_text(i)
        #print(s)
        result += len(s)
    
    print(result)
