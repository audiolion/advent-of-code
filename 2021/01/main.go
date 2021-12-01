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

func run() error {
	in, err := os.ReadFile("input.1")
	if err != nil {
		return err
	}

	lines := strings.Split(string(in), "\n")

	count := 0
	prev, _ := strconv.Atoi(lines[0])
	for i := 1; i < len(lines); i++ {
		curr, _ := strconv.Atoi(lines[i])
		if curr > prev {
			count++
		}
		prev = curr
	}

	fmt.Println(count)

	count3 := 0
	window := []int{0, 0, 0}
	prevSum := 0
	for i := 0; i < len(lines); i++ {
		curr, _ := strconv.Atoi(lines[i])
		if i < 3 {
			window[i] = curr
			prevSum += curr
			continue
		}

		newSum := prevSum - window[0] + curr
		window = append(window[1:], curr)
		if newSum > prevSum {
			count3++
		}
		prevSum = newSum
	}

	fmt.Println(count3)

	return nil
}
