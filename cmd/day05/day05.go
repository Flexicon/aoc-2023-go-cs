package main

import (
	"aoc23/pkg/util"
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	Input string
	//go:embed test_input.txt
	TestInput string
)

type Mapping struct {
	Dst int
	Src int
	Len int
}

func main() {
	fmt.Println("AOC Day 05 🎉")

	util.Check(1, part1, 35, TestInput, Input)
	util.Check(2, part2, 46, TestInput, Input)
}

func part1(input string) int {
	sections := strings.Split(input, "\n\n")
	seeds := parseInts(strings.Split(sections[0], ": ")[1])

	return minLocationForSeeds(sections[1:], seeds)
}

func part2(input string) int {
	sections := strings.Split(input, "\n\n")
	seeds := calcSeedsFromRanges(parseInts(strings.Split(sections[0], ": ")[1]))

	return minLocationForSeeds(sections[1:], seeds)
}

func minLocationForSeeds(sections []string, seeds []int) int {
	values := seeds[:]

	for _, section := range sections {
		mappings := []Mapping{}

		for _, line := range strings.Split(section, "\n")[1:] {
			nums := parseInts(line)
			mappings = append(mappings, Mapping{nums[0], nums[1], nums[2]})
		}

		for i := 0; i < len(values); i++ {
			val := values[i]

			for _, mapping := range mappings {
				src, len, dst := mapping.Src, mapping.Len, mapping.Dst

				if src <= val && src+len > val {
					values[i] = dst + val - src
				}
			}
		}
	}

	return slices.Min(values)
}

func calcSeedsFromRanges(ranges []int) []int {
	seeds := []int{}

	for i := 0; i < len(ranges); i += 2 {
		start, len := ranges[i], ranges[i+1]

		for j := start; j < start+len; j++ {
			seeds = append(seeds, j)
		}
	}

	return seeds
}

func parseInts(raw string) []int {
	vals := strings.Split(strings.TrimSpace(raw), " ")
	ints := make([]int, len(vals))

	for i, v := range vals {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("failed to parse '%s' as int", v)
		}
		ints[i] = n
	}

	return ints
}
