package main

import (
	"aoc23/pkg/util"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	Input string
	//go:embed test_input1.txt
	TestInput1 string
	//go:embed test_input2.txt
	TestInput2 string
)

func main() {
	fmt.Println("AOC Day 01 🎉")

	util.Check(1, part1, 142, TestInput1, Input)
	util.Check(2, part2, 281, TestInput2, Input)
}

func part1(lines []string) int {
	return calculateCalibration(lines, false)
}

func part2(lines []string) int {
	return calculateCalibration(lines, true)
}

var numWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func calculateCalibration(lines []string, handleWords bool) int {
	pat := regexp.MustCompile(`\d`)
	total := 0

	for i, line := range lines {
		if handleWords {
			for j, word := range numWords {
				line = strings.ReplaceAll(line, word, word+strconv.Itoa(j+1)+word)
			}
		}

		found := pat.FindAllString(strings.ToLower(line), -1)
		if len(found) == 0 {
			fmt.Printf("did not find any digits in string on line %d\n", i+1)
			continue
		}

		pair := fmt.Sprintf("%s%s", found[0], found[len(found)-1])
		num, err := strconv.Atoi(pair)
		if err != nil {
			fmt.Printf("failed to parse '%s' as int on line %d\n", pair, i+1)
		}

		total += num
	}

	return total
}
