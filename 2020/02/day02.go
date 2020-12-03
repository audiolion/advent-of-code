package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	if err := part1(); err != nil {
		log.Fatal(err)
	}
	if err := part2(); err != nil {
		log.Fatal(err)
	}
}

func part1() error {
	b, err := ioutil.ReadFile("1.in")
	if err != nil {
		return err
	}

	valid := 0

	contents := string(b)
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		freqs := strings.Split(parts[0], "-")

		low, err := strconv.Atoi(freqs[0])
		if err != nil {
			return err
		}
		high, err := strconv.Atoi(freqs[1])
		if err != nil {
			return err
		}

		ch := strings.TrimSuffix(parts[1], ":")
		pw := parts[2]

		count := strings.Count(pw, ch)
		if count >= low && count <= high {
			valid++
		}
	}

	fmt.Println(valid)
	return nil
}

func part2() error {
	b, err := ioutil.ReadFile("2.in")
	if err != nil {
		return err
	}

	valid := 0

	contents := string(b)
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		indices := strings.Split(parts[0], "-")

		low, err := strconv.Atoi(indices[0])
		if err != nil {
			return err
		}
		high, err := strconv.Atoi(indices[1])
		if err != nil {
			return err
		}

		ch := strings.TrimSuffix(parts[1], ":")
		pw := parts[2]

		count := 0
		if len(pw) >= low && string(pw[low-1]) == ch {
			count++
		}
		if len(pw) >= high && string(pw[high-1]) == ch {
			count++
		}
		if count == 1 {
			valid++
		}
	}

	fmt.Println(valid)
	return nil
}
