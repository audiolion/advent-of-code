package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type Point struct {
	x int
	y int
}

type Vent struct {
	a *Point
	b *Point
}

func run() error {
	b, err := os.ReadFile("1.in")
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")

	vents := make([]*Vent, 0, len(lines))

	//x1,y1 -> x2,y2
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		coord := strings.Split(coords[0], ",")
		x1, err := strconv.Atoi(coord[0])
		if err != nil {
			return err
		}
		y1, err := strconv.Atoi(coord[1])
		if err != nil {
			return err
		}
		a := &Point{x: x1, y: y1}

		coord = strings.Split(coords[1], ",")
		x2, err := strconv.Atoi(coord[0])
		if err != nil {
			return err
		}
		y2, err := strconv.Atoi(coord[1])
		if err != nil {
			return err
		}
		b := &Point{x: x2, y: y2}

		vents = append(vents, &Vent{a: a, b: b})
	}

	xMax := 0
	yMax := 0

	for _, vent := range vents {
		if vent.a.x > xMax {
			xMax = vent.a.x
		}
		if vent.a.y > yMax {
			yMax = vent.a.y
		}
		if vent.b.x > xMax {
			xMax = vent.b.x
		}
		if vent.b.y > yMax {
			yMax = vent.b.y
		}
	}

	xMax++
	yMax++
	fmt.Println(xMax, yMax)

	grid1 := make([][]int, xMax)
	for i := range grid1 {
		grid1[i] = make([]int, yMax)
	}

	grid2 := make([][]int, xMax)
	for i := range grid2 {
		grid2[i] = make([]int, yMax)
	}

	grid1 = mapGridVents(grid1, vents, false)
	grid2 = mapGridVents(grid2, vents, true)

	p1 := 0
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			if grid1[x][y] >= 2 {
				p1++
			}
		}
	}

	p2 := 0
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			if grid2[x][y] >= 2 {
				p2++
			}
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)

	return nil
}

func mapGridVents(grid [][]int, vents []*Vent, mapDiagonals bool) [][]int {
	for _, vent := range vents {
		horizontal := vent.a.x == vent.b.x
		vertical := vent.a.y == vent.b.y
		diagonal := !horizontal && !vertical

		if horizontal {
			x := vent.a.x
			ay := vent.a.y
			by := vent.b.y
			if vent.a.y > vent.b.y {
				ay = vent.b.y
				by = vent.a.y
			}
			for y := ay; y <= by; y++ {
				grid[x][y]++
			}
		}

		if vertical {
			y := vent.a.y
			ax := vent.a.x
			bx := vent.b.x
			if vent.a.x > vent.b.x {
				ax = vent.b.x
				bx = vent.a.x
			}
			for x := ax; x <= bx; x++ {
				grid[x][y]++
			}
		}

		if mapDiagonals && diagonal {
			if vent.a.x > vent.b.x {
				vent.a, vent.b = vent.b, vent.a
			}

			direction := 1
			if vent.a.y > vent.b.y {
				direction = -1
			}

			y := vent.a.y
			for x := vent.a.x; x <= vent.b.x; x++ {
				if y < 0 || y >= len(grid) {
					break
				}
				grid[x][y]++
				y += direction
			}
		}
	}
	return grid
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}
