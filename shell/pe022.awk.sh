#!/bin/sh

# Project Euler, Problem 22:
#
# Begin by sorting the names in input.txt it into alphabetical order.
# Then working out the alphabetical value for each name, multiply this value
# by its alphabetical position in the list to obtain a name score.
# What is the total of all the name scores in the file?
#
# https://projecteuler.net/problem=22

cat pe022_input.txt | sed 's/,/\n/g; s/"//g' | sort | awk -F '' '
BEGIN {
	for (i = 1; i <= 26; i++) {
		char = sprintf("%c", i + 64);
		ord[char] = i;
	}
} 

{
	result = 0
	for (i = 1; i <= NF; i++) {
		result += ord[$i]
	}
	sum += result * NR
}

END {
	print sum
}'

