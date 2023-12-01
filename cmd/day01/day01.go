package main

import (
	_ "embed"
	"fmt"
	"os"
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

	t1Ans := 142
	t1Res := part1(readLines(TestInput1))
	if t1Res != t1Ans {
		fmt.Printf("Test Part 1 -> Expected %d - got %d\n", t1Ans, t1Res)
		os.Exit(1)
	}

	p1Res := part1(readLines(Input))
	fmt.Printf("\nPart 1 result: %d\n\n", p1Res)

	t2Ans := 281
	t2Res := part2(readLines(TestInput2))
	if t2Res != t2Ans {
		fmt.Printf("Test Part 2 -> Expected %d - got %d\n", t2Ans, t2Res)
		os.Exit(1)
	}

	p2Res := part2(readLines(Input))
	fmt.Printf("\nPart 2 result: %d\n\n", p2Res)
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

func readLines(input string) []string {
	return strings.Split(input, "\n")
}
