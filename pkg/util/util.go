package util

import (
	"fmt"
	"os"
)

func Check(part int, fn func(string) int, expected int, testInput, input string) {
	testRes := fn(testInput)
	if testRes != expected {
		fmt.Printf("Test Part %d -> Expected %d - got %d\n", part, expected, testRes)
		os.Exit(1)
	}

	answer := fn(input)
	fmt.Printf("\nPart %d result: %d\n\n", part, answer)
}
