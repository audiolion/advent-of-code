package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const forward = "forward"
const up = "up"
const down = "down"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	b, err := os.ReadFile("input.1")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	horizontal := 0
	depth := 0
	for _, line := range lines {
		xs := strings.SplitN(line, " ", 2)
		command := xs[0]
		magnitude, err := strconv.Atoi(xs[1])
		if err != nil {
			return err
		}

		switch command {
		case forward:
			horizontal += magnitude
		case up:
			depth -= magnitude
		case down:
			depth += magnitude
		default:
			return fmt.Errorf("unrecognized command: %q", command)
		}
	}

	fmt.Println(horizontal * depth)

	horizontal = 0
	depth = 0
	aim := 0
	for _, line := range lines {
		xs := strings.Split(line, " ")
		command := xs[0]
		magnitude, err := strconv.Atoi(xs[1])
		if err != nil {
			return err
		}

		switch command {
		case forward:
			horizontal += magnitude
			depth += magnitude * aim
		case up:
			aim -= magnitude
		case down:
			aim += magnitude
		default:
			return fmt.Errorf("unrecognized command: %q", command)
		}
	}

	fmt.Println(horizontal * depth)

	return nil
}
