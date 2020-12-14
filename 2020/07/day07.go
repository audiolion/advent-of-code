package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("1.in")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	bagMap, err := makeBagMap(lines)
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(bagMap)
	puzzle2(bagMap)
}

func puzzle1(bagMap map[string]bag) {
	ans := 0
	for _, v := range bagMap {
		if leadsToGold(v.contains, bagMap) {
			ans++
		}
	}
	fmt.Println(ans)
}

func leadsToGold(bags []bag, bagMap map[string]bag) bool {
	for _, b := range bags {
		if b.color == "shiny gold" {
			return true
		}
		if leadsToGold(bagMap[b.color].contains, bagMap) {
			return true
		}
	}
	return false
}

func puzzle2(bagMap map[string]bag) {
	ans := countBags(bagMap["shiny gold"], bagMap) - 1
	fmt.Println(ans)
}

func countBags(b bag, bagMap map[string]bag) int {
	count := b.quantity
	if len(bagMap[b.color].contains) == 0 {
		return count
	}
	for _, b2 := range bagMap[b.color].contains {
		count += b.quantity * countBags(b2, bagMap)
	}
	return count
}

type bag struct {
	quantity int
	color    string
	contains []bag
}

func makeBagMap(lines []string) (map[string]bag, error) {
	bagMap := make(map[string]bag)
	r := regexp.MustCompile(`(?P<quanity>[0-9]+) (?P<color>[a-z ]+) bags?`)
	for _, line := range lines {
		split := strings.Split(strings.TrimSuffix(line, "."), " bags contain ")
		color := split[0]

		var newBag bag

		if strings.Contains(split[1], "no other bags") {
			newBag = bag{quantity: 1, color: color, contains: make([]bag, 0)}
			bagMap[newBag.color] = newBag
			continue
		}

		bags := strings.Split(split[1], ", ")
		newBag = bag{quantity: 1, color: color, contains: make([]bag, 0, len(bags))}
		for _, b := range bags {
			res := r.FindStringSubmatch(b)

			q, err := strconv.Atoi(res[1])
			if err != nil {
				return nil, err
			}

			newBag.contains = append(newBag.contains, bag{quantity: q, color: res[2], contains: []bag{}})
		}

		bagMap[newBag.color] = newBag
	}
	return bagMap, nil
}
