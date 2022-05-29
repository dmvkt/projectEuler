/*
Project Euler, Problem 19:

Given that 1 Jan 1900 was a Monday, how many Sundays fell on the first of
the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

https://projecteuler.net/problem=19
*/

package main

import "fmt"

func main() {
	daysInMonth := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	dayCnt, sundayCnt := 0, 0
	for year := 1900; year <= 2000; year += 1 {
		for month := 1; month <= 12; month += 1 {
			if year >= 1901 && dayCnt%7 == 0 {
				sundayCnt += 1
			}
			dayCnt += daysInMonth[month-1]
			if month == 2 && year%4 == 0 {
				dayCnt += 1
			}
		}
	}
	fmt.Println(sundayCnt)
}
