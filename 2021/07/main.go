package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	b, err := os.ReadFile("1.in")
	if err != nil {
		return err
	}

	nums := strings.Split(string(b), ",")

	xs := []int{}

	for _, num := range nums {
		x, err := strconv.Atoi(num)
		if err != nil {
			return err
		}
		xs = append(xs, x)
	}

	max := math.MinInt
	for _, x := range xs {
		if max < x {
			max = x
		}
	}

	positions := make([]int, max)

	minFuel := math.MaxInt
	for candidate := range positions {
		fuel := 0
		for _, x := range xs {
			fuel += abs(candidate - x)
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)

	minFuel = math.MaxInt
	for candidate := range positions {
		fuel := 0
		for _, x := range xs {
			distance := abs(candidate - x)

			// (x + y) * z / 2 where x = lowest number in sequence, y is highest, z is number of items in sequence
			// in our case the sequence is always 1+..+distance
			additionalFuel := (1 + distance) * distance / 2

			fuel += additionalFuel
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)

	return nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
