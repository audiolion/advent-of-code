package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("1.in")
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(b), "\n")

	puzzle1(instructions)
	puzzle2(instructions)
}

func puzzle1(instructions []string) {
	_, acc := isInfinite(instructions)
	fmt.Println(acc)
}

func puzzle2(instructions []string) {
	for i, line := range instructions {
		tmp := make([]string, len(instructions))
		copy(tmp, instructions)

		if strings.Contains(line, "nop") {
			tmp[i] = strings.Replace(line, "nop", "jmp", 1)
			inf, acc := isInfinite(tmp)
			if !inf {
				fmt.Println(acc)
				return
			}
		}

		if strings.Contains(line, "jmp") {
			tmp[i] = strings.Replace(line, "jmp", "nop", 1)
			inf, acc := isInfinite(tmp)
			if !inf {
				fmt.Println(acc)
				return
			}
		}
	}
}

func isInfinite(instructions []string) (bool, int) {
	i := 0
	acc := 0
	seen := make(map[int]bool)
	for {
		if i == len(instructions) {
			return false, acc
		}

		if seen[i] {
			return true, acc
		}
		seen[i] = true

		op, num, err := parseOpcode(instructions[i])
		if err != nil {
			log.Fatal(err)
		}

		switch op {
		case "nop":
			i++
		case "acc":
			acc += num
			i++
		case "jmp":
			i += num
		}
	}
}

func parseOpcode(instruction string) (op string, num int, err error) {
	opcode := strings.Split(instruction, " ")
	op = opcode[0]
	num, err = strconv.Atoi(opcode[1])
	return
}
