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
	b, err := os.ReadFile("1.in")
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")

	numbers := strings.Split(lines[0], ",")
	nums := make([]int, len(numbers))
	for i, ch := range numbers {
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return err
		}
		nums[i] = num
	}

	boardLines := lines[2:]
	boards := make([][5][5]int, 0, len(boardLines)/5)
	board := [5][5]int{}
	row := 0
	for _, line := range boardLines {
		if len(line) == 0 {
			boards = append(boards, board)
			board = [5][5]int{}
			row = 0
			continue
		}

		line = strings.ReplaceAll(strings.TrimSpace(line), "  ", " ")
		chars := strings.Split(line, " ")
		for i, ch := range chars {
			num, err := strconv.Atoi(ch)
			if err != nil {
				return err
			}
			board[row][i] = num
		}
		row++
	}
	boards = append(boards, board)

	p1Boards := boards[:]

	var winner [5][5]int
	var winningNumber int

loop:
	for _, num := range nums {
		for boardIdx, board := range p1Boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board[i][j] == num {
						board[i][j] = -1
					}
				}
			}
			p1Boards[boardIdx] = board
		}
		for _, board := range p1Boards {
			for i := 0; i < 5; i++ {
				rowBingo := 0
				colBingo := 0
				for j := 0; j < 5; j++ {
					rowBingo += board[i][j]
					colBingo += board[j][i]
				}
				if rowBingo == -5 || colBingo == -5 {
					winner = board
					winningNumber = num
					break loop
				}
			}
		}
	}

	sum := 0
	for _, row := range winner {
		for _, n := range row {
			if n != -1 {
				sum += n
			}
		}
	}

	fmt.Println(sum * winningNumber)

	p2Boards := boards[:]
loop2:
	for len(p2Boards) != 1 {
		for _, num := range nums {
			for boardIdx, board := range p2Boards {
				for i := 0; i < 5; i++ {
					for j := 0; j < 5; j++ {
						if board[i][j] == num {
							board[i][j] = -1
						}
					}
				}
				p2Boards[boardIdx] = board
			}
			winners := make(map[int]struct{})
			for boardIdx, board := range p2Boards {
				for i := 0; i < 5; i++ {
					rowBingo := 0
					colBingo := 0
					for j := 0; j < 5; j++ {
						rowBingo += board[i][j]
						colBingo += board[j][i]
					}
					if rowBingo == -5 || colBingo == -5 {
						winners[boardIdx] = struct{}{}
					}
				}
			}
			newBoards := make([][5][5]int, 0, len(p2Boards))
			for boardIdx, board := range p2Boards {
				_, winner := winners[boardIdx]
				if !winner {
					newBoards = append(newBoards, board)
				}
			}
			if len(p2Boards) == 1 && len(newBoards) == 0 {
				winningNumber = num
				winner = p2Boards[0]
				break loop2
			}
			p2Boards = newBoards
		}
	}

	fmt.Println(winningNumber)
	for _, row := range winner {
		fmt.Println(row)
	}

	sum = 0
	for _, row := range winner {
		for _, n := range row {
			if n != -1 {
				sum += n
			}
		}
	}

	fmt.Println(sum * winningNumber)

	return nil
}
