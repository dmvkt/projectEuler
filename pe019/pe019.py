#!/usr/bin/python

"""
Project Euler, Problem 19:

Given that 1 Jan 1900 was a Monday, how many Sundays fell on the first of
the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

https://projecteuler.net/problem=19
"""

def main():
    days_in_month = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    year, month = 1900, 1
    day_cnt, sunday_cnt = 0, 0
    while year <= 2000:
        while month <= 12:
            if year >= 1901 and day_cnt % 7 == 0:
                sunday_cnt += 1
            day_cnt += days_in_month[month - 1]
            if month == 2 and year % 4 == 0:
                day_cnt += 1
            month += 1
        year += 1
        month = 1
    print(sunday_cnt)

if __name__ == "__main__":
    main()
