package main

import (
	"aoc23/pkg/util"
	_ "embed"
	"fmt"
	"log"
	"math"
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
	seedRanges := buildSeedRanges(parseInts(strings.Split(sections[0], ": ")[1]))

	min := math.MaxInt
	for i, sr := range seedRanges {
		chunks := chunked(sr, 1_000_000)
		count := len(chunks)
		fmt.Printf("#%d - Chunked Ranges for (%d,%d): %+v\n", i+1, sr[0], sr[1], count)

		for j, chunk := range chunks {
			pct := int((float64(j+1) / float64(count)) * 100)
			fmt.Printf("#%d\tchecking range: %+v [%d/%d] %d%%\n", i+1, chunk, j+1, count, pct)

			if n := minLocationForSeeds(sections, genSeedsFromRange(chunk)); n < min {
				min = n
			}
		}
		fmt.Println()
	}

	return min
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

func buildSeedRanges(nums []int) [][]int {
	ranges := [][]int{}

	for i := 0; i < len(nums); i += 2 {
		start, len := nums[i], nums[i+1]
		ranges = append(ranges, []int{start, start + len - 1})
	}

	return ranges
}

func chunked(r []int, size int) [][]int {
	from, to := r[0], r[1]

	if to-from <= size {
		return [][]int{r}
	}

	next := []int{from, from + size}
	chunks := [][]int{next}

	return append(chunks, chunked([]int{from + size + 1, to}, size)...)
}

func genSeedsFromRange(r []int) []int {
	from, to := r[0], r[1]
	seeds := []int{}

	for i := from; i <= to; i++ {
		seeds = append(seeds, i)
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
