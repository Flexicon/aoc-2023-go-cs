package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	Input string
	//go:embed test_input.txt
	TestInput string
)

func main() {
	fmt.Println("AOC Day 02 🎉")

	check(1, part1, TestInput, 8)
	check(2, part2, TestInput, 2286)
}

func part1(lines []string) int {
	maxColors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0

	for i, line := range lines {
		gameID := i + 1
		rounds := strings.Split(strings.Split(line, ": ")[1], "; ")
		isValid := true

		for _, round := range rounds {
			sets := strings.Split(round, ", ")

			for _, set := range sets {
				pair := strings.Split(set, " ")
				color := pair[1]
				amnt, err := strconv.Atoi(pair[0])
				if err != nil {
					panic(fmt.Sprintf("could not parse '%s' as int", pair[0]))
				}

				if amnt > maxColors[color] {
					isValid = false
				}
			}
		}

		if isValid {
			sum += gameID
		}
	}

	return sum
}

func part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		game := strings.Split(strings.Split(line, ": ")[1], "; ")
		mins := map[string]int{"red": 0, "blue": 0, "green": 0}

		for _, round := range game {
			sets := strings.Split(round, ", ")

			for _, set := range sets {
				pair := strings.Split(set, " ")
				color := pair[1]
				amnt, err := strconv.Atoi(pair[0])
				if err != nil {
					panic(fmt.Sprintf("could not parse '%s' as int", pair[0]))
				}

				if mins[color] < amnt {
					mins[color] = amnt
				}
			}
		}

		sum += mins["red"] * mins["blue"] * mins["green"]
	}

	return sum
}

func check(part int, fn func([]string) int, testInput string, expected int) {
	testRes := fn(strings.Split(testInput, "\n"))
	if testRes != expected {
		fmt.Printf("Test Part %d -> Expected %d - got %d\n", part, expected, testRes)
		os.Exit(1)
	}

	answer := fn(strings.Split(Input, "\n"))
	fmt.Printf("\nPart %d result: %d\n\n", part, answer)
}
