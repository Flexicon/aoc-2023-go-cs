package main

import (
	"aoc23/pkg/util"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	//go:embed input.txt
	Input string
	//go:embed test_input.txt
	TestInput string
)

func main() {
	fmt.Println("AOC Day 03 🎉")

	util.Check(1, part1, 4361, TestInput, Input)
	util.Check(2, part2, 467835, TestInput, Input)
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for i, line := range lines {
		chars := []rune(line)
		cur := strings.Builder{}
		isPart := false

		for j, c := range chars {
			if !unicode.IsDigit(c) {
				sum += calcPartNumber(cur, isPart)
				cur.Reset()
				isPart = false
				continue
			}
			cur.WriteRune(c)

			for _, neighbor := range getNeighbors(i, j, lines) {
				if isSym(neighbor.Value) {
					isPart = true
					break
				}
			}
		}

		sum += calcPartNumber(cur, isPart)
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for i, line := range lines {
		chars := []rune(line)

		for j, c := range chars {
			if c != '*' {
				continue
			}

			checked := map[string]bool{}
			neighboringParts := []int{}

			for _, neighbor := range getNeighbors(i, j, lines) {
				var x, y = neighbor.X, neighbor.Y
				if checked[fmt.Sprintf("%d,%d", x, y)] {
					continue
				}

				if unicode.IsDigit(neighbor.Value) {
					checked[fmt.Sprintf("%d,%d", x, y)] = true
					neighborLine := []rune(lines[x])
					tmp := string(neighbor.Value)

					for z := y - 1; z >= 0; z-- {
						checked[fmt.Sprintf("%d,%d", x, z)] = true
						if !unicode.IsDigit(neighborLine[z]) {
							break
						}
						tmp = string(neighborLine[z]) + tmp
					}

					for z := y + 1; z < len(neighborLine); z++ {
						checked[fmt.Sprintf("%d,%d", x, z)] = true
						if !unicode.IsDigit(neighborLine[z]) {
							break
						}
						tmp += string(neighborLine[z])
					}

					partNum, err := strconv.Atoi(tmp)
					if err != nil {
						panic(fmt.Sprintf("invalid part number: %s", tmp))
					}
					neighboringParts = append(neighboringParts, partNum)
				}
			}

			if len(neighboringParts) > 1 {
				ratio := 1
				for _, p := range neighboringParts {
					ratio *= p
				}
				sum += ratio
			}
		}
	}

	return sum
}

func isSym(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
}

func calcPartNumber(pb strings.Builder, isPart bool) int {
	if !isPart || pb.Len() == 0 {
		return 0
	}

	partNum, err := strconv.Atoi(pb.String())
	if err != nil {
		panic(fmt.Sprintf("invalid part number: %s", pb.String()))
	}
	return partNum
}

type Neighbor struct {
	Value rune
	X     int
	Y     int
}

func getNeighbors(i, j int, lines []string) []Neighbor {
	return []Neighbor{
		getNeighborCell(i-1, j, lines),
		getNeighborCell(i-1, j+1, lines),
		getNeighborCell(i, j+1, lines),
		getNeighborCell(i+1, j+1, lines),
		getNeighborCell(i+1, j, lines),
		getNeighborCell(i+1, j-1, lines),
		getNeighborCell(i, j-1, lines),
		getNeighborCell(i-1, j-1, lines),
	}
}

func getNeighborCell(i, j int, lines []string) Neighbor {
	value := '.'
	if i >= 0 && j >= 0 && i < len(lines) && j < len(lines[i]) {
		value = []rune(lines[i])[j]
	}
	return Neighbor{value, i, j}
}
