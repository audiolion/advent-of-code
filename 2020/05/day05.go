package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("1.in")
	if err != nil {
		log.Fatal(err)
	}
	contents := string(b)

	if err := puzzle1(contents); err != nil {
		log.Fatal(err)
	}
	if err := puzzle2(contents); err != nil {
		log.Fatal(err)
	}
}

type direction int

const (
	down direction = iota
	up
)

func partition(d direction, lo, hi int) (int, int) {
	if d == down {
		return lo, lo + (hi-lo)/2
	}
	if d == up {
		return lo + (hi-lo)/2 + 1, hi
	}
	return -1, -1
}

func findSeat(rowIn, colIn string) (row int, col int) {
	var (
		rowLo int = 0
		colLo     = 0
		rowHi     = 127
		colHi     = 7
	)

	for _, ch := range rowIn {
		if ch == 'F' {
			rowLo, rowHi = partition(down, rowLo, rowHi)

		} else if ch == 'B' {
			rowLo, rowHi = partition(up, rowLo, rowHi)
		}
		if rowLo == rowHi {
			row = rowLo
			break
		}
	}
	for _, ch := range colIn {
		if ch == 'L' {
			colLo, colHi = partition(down, colLo, colHi)
		} else if ch == 'R' {
			colLo, colHi = partition(up, colLo, colHi)
		}
		if colLo == colHi {
			col = colLo
			break
		}
	}
	return row, col
}

func puzzle1(input string) error {
	max := -1

	for _, line := range strings.Split(input, "\n") {
		row, col := findSeat(line[:7], line[7:])
		seatID := row*8 + col
		if seatID > max {
			max = seatID
		}
	}

	fmt.Println(max)
	return nil
}

func puzzle2(input string) error {
	seatIDs := make(map[int]bool)

	for _, line := range strings.Split(input, "\n") {
		row, col := findSeat(line[:7], line[7:])
		seatID := row*8 + col
		seatIDs[seatID] = true
	}

	for k := range seatIDs {
		if !seatIDs[k+1] && seatIDs[k+2] {
			fmt.Println(k + 1)
			break
		}
		if !seatIDs[k-1] && seatIDs[k-2] {
			fmt.Println(k - 1)
			break
		}
	}
	return nil
}
