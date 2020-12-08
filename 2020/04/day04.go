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
	if err := puzzle1(); err != nil {
		log.Fatal(err)
	}
	if err := puzzle2(); err != nil {
		log.Fatal(err)
	}
}

func puzzle1() error {
	b, err := ioutil.ReadFile("1.in")
	if err != nil {
		return err
	}
	contents := string(b)
	valid := 0
	passports := strings.Split(contents, "\n\n")
	for _, p := range passports {
		if strings.Contains(p, "byr") && strings.Contains(p, "iyr") && strings.Contains(p, "eyr") && strings.Contains(p, "hgt") && strings.Contains(p, "hcl") && strings.Contains(p, "ecl") && strings.Contains(p, "pid") {
			valid++
		}
	}
	fmt.Println(valid)
	return nil
}

func puzzle2() error {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.
	b, err := ioutil.ReadFile("1.in")
	if err != nil {
		return err
	}
	contents := string(b)
	validCount := 0
	passports := strings.Split(contents, "\n\n")
	for _, p := range passports {
		valid := true
		if !strings.Contains(p, "byr") || !strings.Contains(p, "iyr") || !strings.Contains(p, "eyr") || !strings.Contains(p, "hgt") || !strings.Contains(p, "hcl") || !strings.Contains(p, "ecl") || !strings.Contains(p, "pid") {
			valid = false
			continue
		}
		for _, line := range strings.Split(p, "\n") {
			if !valid {
				break
			}
			for _, field := range strings.Split(line, " ") {
				vs := strings.Split(field, ":")
				if vs[0] != "byr" && vs[0] != "iyr" && vs[0] != "eyr" && vs[0] != "hgt" && vs[0] != "hcl" && vs[0] != "ecl" && vs[0] != "pid" && vs[0] != "cid" {
					valid = false
					break
				}
				if len(vs) != 2 {
					valid = false
					break
				}
				if vs[0] == "byr" {
					v, err := strconv.Atoi(vs[1])
					if err != nil || v < 1020 || v > 2002 {
						valid = false
						break
					}
				}
				if vs[0] == "iyr" {
					v, err := strconv.Atoi(vs[1])
					if err != nil || v < 2010 || v > 2020 {
						valid = false
						break
					}
				}
				if vs[0] == "eyr" {
					v, err := strconv.Atoi(vs[1])
					if err != nil || v < 2020 || v > 2030 {
						valid = false
						break
					}
				}
				if vs[0] == "hgt" {
					if strings.Contains(vs[1], "cm") {
						height, err := strconv.Atoi(strings.TrimSuffix(vs[1], "cm"))
						if err != nil || height < 150 || height > 193 {
							valid = false
							break
						}
					} else if strings.Contains(vs[1], "in") {
						height, err := strconv.Atoi(strings.TrimSuffix(vs[1], "in"))
						if err != nil || height < 59 || height > 76 {
							valid = false
							break
						}
					} else {
						valid = false
						break
					}
				}
				if vs[0] == "hcl" {
					if len(vs[1]) != 7 {
						valid = false
						break
					} else {
						match, err := regexp.Match("^#[0-9a-f]+$", []byte(vs[1]))
						if !match || err != nil {
							valid = false
							break
						}
					}
				}
				if vs[0] == "ecl" {
					match, err := regexp.Match("^(amb|blu|brn|gry|grn|hzl|oth)$", []byte(vs[1]))
					if !match || err != nil {
						valid = false
						break
					}
				}
				if vs[0] == "pid" {
					if len(vs[1]) != 9 {
						valid = false
						break
					} else {
						match, err := regexp.Match("^[0-9]+$", []byte(vs[1]))
						if !match || err != nil {
							valid = false
							break
						}
					}
				}
			}
		}
		if valid {
			validCount++
		}
	}
	fmt.Println(validCount)
	return nil
}
