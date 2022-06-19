#!/bin/sh

# Project Euler, Problem 1:
#
# Find the sum of all the multiples of 3 or 5 below 1000.
#
# https://projecteuler.net/problem=1

DEFAULT=1000

print_usage() {
	echo "Usage: $(basename $0) [N=$DEFAULT]"
	echo "Find the sum of all the multiples of 3 or 5 below N."
}

[ "$1" = "--help" -o "$1" = "-h" ] && print_usage && exit

N=${1:-$DEFAULT}

sum=0
i=2
while true; do
	i=$((i + 1))
	[ "$i" -ge "$N" ] && break
	if [ "$((i % 3))" = 0 ] || [ "$((i % 5))" = 0 ]; then
		sum=$((sum + i))
	fi
done
echo $sum
