#!/bin/sh

# Project Euler, Problem 2:
#
# By considering the terms in the Fibonacci sequence whose values do not exceed
# four million, find the sum of the even-valued terms.
#
# https://projecteuler.net/problem=2

DEFAULT=4000000

print_usage() {
	echo "Usage: $(basename $0) [N=$DEFAULT]"
	echo "Find the sum of the even-valued Fibonacci sequence " \
	     "terms, whose values do not exceed N."
}

[ "$1" = "--help" -o "$1" = "-h" ] && print_usage && exit

N=${1:-$DEFAULT}

fib0=1
fib1=1
fib2=0
sum=0
while true; do
	fib2=$((fib0 + fib1))
	[ "$fib2" -ge "$N" ] && break
	[ "$((fib2%2))" = 0 ] && sum=$((sum + fib2))
	fib0=$fib1
	fib1=$fib2
done
echo $sum
