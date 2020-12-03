package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	if err := puzzle1(); err != nil {
		log.Fatal(err)
	}
	if err := puzzle2(); err != nil {
		log.Fatal(err)
	}
}

func puzzle1() error {
	b, err := ioutil.ReadFile("1.in")
	if err != nil {
		return err
	}

	contents := string(b)
	lines := strings.Split(contents, "\n")

	trees := make([][]bool, len(lines))
	for i, line := range lines {
		trees[i] = make([]bool, len(line))
		for j, ch := range line {
			if ch == '#' {
				trees[i][j] = true
			}
		}
	}

	fmt.Println(treeCount(trees, 3, 1))
	return nil
}

func treeCount(trees [][]bool, xSlope, ySlope int) int {
	count := 0
	x := 0
	for row := 0; row < len(trees)-ySlope; row += ySlope {
		if trees[row+ySlope][(x+xSlope)%len(trees[row])] {
			count++
		}
		x += xSlope
	}
	return count
}

func puzzle2() error {
	b, err := ioutil.ReadFile("1.in")
	if err != nil {
		return err
	}

	contents := string(b)
	lines := strings.Split(contents, "\n")

	trees := make([][]bool, len(lines))
	for i, line := range lines {
		trees[i] = make([]bool, len(line))
		for j, ch := range line {
			if ch == '#' {
				trees[i][j] = true
			}
		}
	}

	type slope struct {
		x, y int
	}
	slopes := []slope{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}

	total := 1
	for _, slope := range slopes {
		total *= treeCount(trees, slope.x, slope.y)
	}
	fmt.Println(total)
	return nil
}
