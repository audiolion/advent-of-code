package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Brute force solution

func main() {
	if err := run1(); err != nil {
		log.Fatal(err)
	}
	if err := run2(); err != nil {
		log.Fatal(err)
	}
}

func run1() error {
	f, err := os.Open("input-1.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	input := strings.Split(string(data), "\n")

	prev := make([]int, len(input), len(input))
	for i, s := range input {
		x, err := strconv.Atoi(s)
		if err != nil && len(s) != 0 {
			return err
		}
		for _, y := range prev {
			if x+y == 2020 {
				fmt.Printf("1: %d + %d = 2020; %d * %d = %d\n", x, y, x, y, x*y)
			}
		}
		prev[i] = x
	}
	return nil
}

func run2() error {
	f, err := os.Open("input-2.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	input := strings.Split(string(data), "\n")

	prev := make([]int, len(input), len(input))
	for i, s := range input {
		x, err := strconv.Atoi(s)
		if err != nil && len(s) != 0 {
			return err
		}
		for j, y := range prev {
			for k := j + 1; k < len(prev); k++ {
				z := prev[k]
				if x+y+z == 2020 && x > 0 && y > 0 && z > 0 {
					fmt.Printf("2: %d + %d + %d = 2020; %d * %d * %d = %d\n", x, y, z, x, y, z, x*y*z)
				}
			}
		}
		prev[i] = x
	}
	return nil
}
