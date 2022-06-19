#!/bin/sh

# Project Euler, Problem 3:
#
# What is the largest prime factor of the number 600851475143?
#
# https://projecteuler.net/problem=3

DEFAULT=600851475143

print_usage() {
    echo "Usage: $(basename $0) [N=$DEFAULT]"
    echo "Find the largest prime factor of the number N."
}

[ "$1" = "--help" -o "$1" = "-h" ] && print_usage && exit

N=${1:-$DEFAULT}

factor $N | sed 's/^.* //'

