/*
Project Euler, Problem 22:

Begin by sorting the names in input.txt into alphabetical order.
Then working out the alphabetical value for each name, multiply this value
by its alphabetical position in the list to obtain a name score.
What is the total of all the name scores in the file?

https://projecteuler.net/problem=22
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getScore(name string) int {
	score := 0
	for _, c := range name {
		score += int(c) - int('A') + 1
	}
	return score
}

func main() {
	var names []string

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		names = strings.Split(strings.Replace(input.Text(), "\"", "",
			len(input.Text())), ",")
	}
	f.Close()

	sort.Strings(names)

	result := 0
	for i, name := range names {
		result += (i + 1) * getScore(name)
	}
	fmt.Println(result)
}
