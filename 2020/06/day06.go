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

	groups := strings.Split(string(b), "\n\n")

	if err := puzzle1(groups); err != nil {
		log.Fatal(err)
	}
	if err := puzzle2(groups); err != nil {
		log.Fatal(err)
	}
}

func puzzle1(groups []string) error {
	ans := 0
	for _, group := range groups {
		set := make(map[rune]bool)
		for _, ch := range strings.ReplaceAll(group, "\n", "") {
			if !set[ch] {
				set[ch] = true
				ans++
			}
		}
	}

	fmt.Println(ans)
	return nil
}

func puzzle2(groups []string) error {
	ans := 0
	for _, group := range groups {
		set := make(map[rune]int)
		lines := strings.Split(group, "\n")
		for _, line := range lines {
			for _, ch := range line {
				_, ok := set[ch]
				if !ok {
					set[ch] = 1
				} else {
					set[ch]++
				}
				if set[ch] == len(lines) {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
	return nil
}
