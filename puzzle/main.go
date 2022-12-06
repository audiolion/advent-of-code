package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var day = flag.String("day", "", "The puzzle day to download.")
var year = flag.String("year", strconv.Itoa(time.Now().Year()), "The Advent of Code year.")
var output = flag.String("output", "", "Where to write the puzzle input.")
var session = flag.String("session", os.Getenv("AOC_SESSION_TOKEN"), "Your AoC session token, can be set through AOC_SESSION_TOKEN environment variable.")

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()

	var (
		dayNum  int
		yearNum int
		err     error
	)

	if dayNum, err = strconv.Atoi(*day); err != nil {
		return err
	}
	if yearNum, err = strconv.Atoi(*year); err != nil {
		return err
	}
	if len(*session) == 0 {
		return errors.New("missing session token")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", yearNum, dayNum)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: *session,
	})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	puzzle, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	puzzle = []byte(strings.TrimRight(string(puzzle), "\n"))

	if len(*output) == 0 {
		fmt.Fprintf(os.Stdout, "%s", puzzle)
	} else {
		fi, err := os.Create(*output)
		if err != nil {
			if fi != nil {
				defer func() {
					fi.Close()
					os.Remove(fi.Name())
				}()
			}
			return err
		}
		fi.Write(puzzle)
		defer func() {
			if err := fi.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
			}
		}()
	}

	return nil
}
