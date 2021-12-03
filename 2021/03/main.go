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

	bitLen := len(lines[0])

	bytes := []uint16{}
	for _, line := range lines {
		n, err := strconv.ParseUint(line, 2, bitLen)
		if err != nil {
			return err
		}
		bytes = append(bytes, uint16(n))
	}

	counts := make(map[int][]int)
	for i := 0; i < bitLen; i++ {
		counts[i] = []int{0, 0}
	}

	for _, bits := range bytes {
		n := bits
		for i := bitLen - 1; i >= 0; i-- {
			v, _ := counts[i]
			if n%2 == 0 {
				v[0] += 1
			} else {
				v[1] += 1
			}
			counts[i] = v
			n /= 2
		}
	}

	// part 1
	gammaBits := make([]int, bitLen)
	epsilonBits := make([]int, bitLen)
	for k, v := range counts {
		if v[0] > v[1] {
			gammaBits[k] = 0
			epsilonBits[k] = 1
		} else {
			gammaBits[k] = 1
			epsilonBits[k] = 0
		}
	}

	pos := bitLen - 1
	gamma := 0
	epsilon := 0
	for i := 0; i < bitLen; i++ {
		gammaBit := gammaBits[i]
		epsilonBit := epsilonBits[i]
		if gammaBit == 1 {
			gamma |= (1 << pos)
		}
		if epsilonBit == 1 {
			epsilon |= (1 << pos)
		}
		pos--
	}

	fmt.Println(gamma * epsilon)

	// part 2
	oxygenGeneratorBytes := bytes[:]
	for pos := bitLen - 1; pos >= 0; pos-- {
		mask := uint16(1 << pos)
		count := []int{0, 0}
		for _, bits := range oxygenGeneratorBytes {
			msb := bits & mask
			if msb == 0 {
				count[0]++
			} else {
				count[1]++
			}
		}
		keepMsb := uint16(0)
		if count[1] >= count[0] {
			keepMsb = uint16(1)
		}
		newOxyBytes := make([]uint16, 0, len(oxygenGeneratorBytes))
		for _, bits := range oxygenGeneratorBytes {
			msb := (bits & mask) >> pos
			if msb == keepMsb {
				newOxyBytes = append(newOxyBytes, bits)
			}
		}
		oxygenGeneratorBytes = newOxyBytes[:]
		if len(oxygenGeneratorBytes) == 1 {
			break
		}
	}

	cO2ScrubberBytes := bytes[:]
	for pos := bitLen - 1; pos >= 0; pos-- {
		mask := uint16(1 << pos)
		count := []int{0, 0}
		for _, bits := range cO2ScrubberBytes {
			msb := bits & mask
			if msb == 0 {
				count[0]++
			} else {
				count[1]++
			}
		}
		keepMsb := uint16(0)
		if count[1] < count[0] {
			keepMsb = uint16(1)
		}
		newCO2Bytes := make([]uint16, 0, len(cO2ScrubberBytes))
		for _, bits := range cO2ScrubberBytes {
			msb := (bits & mask) >> pos
			if msb == keepMsb {
				newCO2Bytes = append(newCO2Bytes, bits)
			}
		}
		cO2ScrubberBytes = newCO2Bytes[:]
		if len(cO2ScrubberBytes) == 1 {
			break
		}
	}

	fmt.Println(int(oxygenGeneratorBytes[0]) * int(cO2ScrubberBytes[0]))

	return nil
}
