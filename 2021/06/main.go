package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const TOTAL_EXPERIMENT_DAYS = 80
const TOTAL_OCEAN_TAKEOVER_DAYS = 256
const MAX_SIZE = 9
const FISH_REPRODUCE_DAYS = 6
const NEW_FISH_REPRODUCE_DAYS = 8

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

	fish := [MAX_SIZE]int{}

	for _, num := range nums {
		days, err := strconv.Atoi(num)
		if err != nil {
			return err
		}
		fish[days]++
	}

	experiment := runExperiment(fish, TOTAL_EXPERIMENT_DAYS)

	p1 := 0
	for _, qty := range experiment {
		p1 += qty
	}

	fmt.Println(p1)

	takeover := runExperiment(fish, TOTAL_OCEAN_TAKEOVER_DAYS)

	p2 := 0
	for _, qty := range takeover {
		p2 += qty
	}

	fmt.Println(p2)

	return nil
}

func runExperiment(fish [MAX_SIZE]int, numDays int) [MAX_SIZE]int {
	for day := 0; day < numDays; day++ {
		newFish := [MAX_SIZE]int{}
		reproducingFish := fish[0]
		for i := 1; i < len(fish); i++ {
			newFish[i-1] = fish[i]
		}
		newFish[FISH_REPRODUCE_DAYS] += reproducingFish
		newFish[NEW_FISH_REPRODUCE_DAYS] += reproducingFish
		fish = newFish
	}
	return fish
}
