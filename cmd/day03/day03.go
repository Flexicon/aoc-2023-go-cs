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
	util.Check(2, part2, 9999, TestInput, Input)
}

func part1(lines []string) int {
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
				if isSym(neighbor) {
					isPart = true
					break
				}
			}
		}

		sum += calcPartNumber(cur, isPart)
	}

	return sum
}

func part2(lines []string) int {
	return 0
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

func getNeighbors(i, j int, lines []string) []rune {
	return []rune{
		getCell(i-1, j, lines),
		getCell(i-1, j+1, lines),
		getCell(i, j+1, lines),
		getCell(i+1, j+1, lines),
		getCell(i+1, j, lines),
		getCell(i+1, j-1, lines),
		getCell(i, j-1, lines),
		getCell(i-1, j-1, lines),
	}
}

func getCell(i, j int, lines []string) rune {
	if i >= 0 && j >= 0 && i < len(lines) && j < len(lines[i]) {
		return []rune(lines[i])[j]
	}
	return '.'
}
