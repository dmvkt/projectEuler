/*
Project Euler, Problem 11:

What is the greatest product of four adjacent numbers in the same direction
(up, down, left, right, or diagonally) in the 20×20 grid (input.txt)?

https://projecteuler.net/problem=11
*/

/*
Implementation note:
The solution could have been more compact and straightforward, but the goal was
to take an iterator approach.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type iterMode int

const (
	IterOverRow iterMode = iota
	IterOverColumn
	IterOverDiagonal
	IterOverSDiagonal
)

/* Iterator for a 2D array (data) */
type Iter struct {
	data   *[][]int
	i, j   int
	k      int
	onEdge bool
	mode   iterMode
}

func NewIter(data *[][]int) *Iter {
	iter := new(Iter)
	iter.Reset()
	iter.data = data
	return iter
}

func (iter *Iter) Reset() {
	iter.onEdge = false
	if iter.mode == IterOverRow {
		iter.i, iter.j = 0, -1
	} else if iter.mode == IterOverColumn {
		iter.i, iter.j = -1, 0
	} else if iter.mode == IterOverDiagonal {
		iter.k = -len(*iter.data)
		iter.i, iter.j = -iter.k, 0
	} else {
		iter.k = -len(*iter.data)
		iter.i, iter.j = -iter.k, len((*iter.data)[0])
	}
}

func (iter *Iter) Next() bool {
	switch iter.mode {
	case IterOverRow:
		return iter.NextInRow()
	case IterOverColumn:
		return iter.NextInColumn()
	case IterOverDiagonal:
		return iter.NextInDiagonal()
	case IterOverSDiagonal:
		return iter.NextInSDiagonal()
	}
	return false
}

func (iter *Iter) Get() int {
	if iter.onEdge {
		return 0
	}
	return (*iter.data)[iter.i][iter.j]
}

/* Iterate over rows */
func (iter *Iter) NextInRow() bool {
	if iter.onEdge == true {
		iter.onEdge = false
		return true
	}
	iter.j += 1
	if iter.j == len((*iter.data)[0]) {
		iter.j = 0
		iter.i += 1
		iter.onEdge = true
	}
	return iter.i != len(*iter.data)
}

/* Iterate over columns */
func (iter *Iter) NextInColumn() bool {
	if iter.onEdge == true {
		iter.onEdge = false
		return true
	}
	iter.i += 1
	if iter.i == len(*iter.data) {
		iter.i = 0
		iter.j += 1
		iter.onEdge = true
	}
	return iter.j != len((*iter.data)[0])
}

/* Iterate over diagonals */
func (iter *Iter) NextInDiagonal() bool {
	if iter.onEdge == true {
		iter.onEdge = false
		return true
	}
	iter.i += 1
	iter.j += 1
	if iter.i >= len(*iter.data) || iter.j >= len((*iter.data)[0]) {
		iter.k += 1
		if iter.k < 0 {
			iter.i = -iter.k
			iter.j = 0
		} else {
			iter.i = 0
			iter.j = iter.k
		}
		iter.onEdge = true
	}
	return iter.k < len((*iter.data)[0])
}

/* Iterate over secondary diagonals */
func (iter *Iter) NextInSDiagonal() bool {
	if iter.onEdge == true {
		iter.onEdge = false
		return true
	}
	iter.i += 1
	iter.j -= 1
	if iter.i >= len(*iter.data) || iter.j < 0 {
		iter.k += 1
		if iter.k < 0 {
			iter.i = -iter.k
			iter.j = len((*iter.data)[0]) - 1
		} else {
			iter.i = 0
			iter.j = len((*iter.data)[0]) - iter.k - 1
		}
		iter.onEdge = true
	}
	return iter.k < len((*iter.data)[0])
}

func assertNoError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	const n = 4
	var data [][]int
	f, err := os.Open("input.txt")
	assertNoError(err)
	input := bufio.NewScanner(f)

	for input.Scan() {
		var values []int
		if strings.TrimSpace(input.Text()) == "" {
			continue
		}
		for _, s := range strings.Split(input.Text(), " ") {
			value, err := strconv.Atoi(s)
			assertNoError(err)
			values = append(values, value)
		}
		data = append(data, values)
	}
	f.Close()

	maxProduct := 0
	iterator := NewIter(&data)
	for _, mode := range []iterMode{IterOverRow, IterOverColumn, IterOverDiagonal, IterOverSDiagonal} {
		iterator.mode = mode
		iterator.Reset()

		product := 1
		var have []int
		for iterator.Next() {
			value := iterator.Get()
			if value == 0 {
				have = have[:0]
				product = 1
				continue
			}
			product *= value
			have = append(have, value)
			if len(have) < n {
				continue
			}
			if len(have) > n {
				product = int(product / have[len(have)-n-1])
			}
			if product > maxProduct {
				maxProduct = product
			}
		}
	}
	fmt.Println(maxProduct)
}
