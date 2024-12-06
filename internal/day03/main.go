package day03

import (
	"github.com/lewis-od/aoc24/internal/day03/parsing"
)

func Part1(input string) int {
	scanner := parsing.NewScanner(input)
	parser := parsing.NewParser(scanner.ScanTokens())
	operations := parser.Parse()

	result := 0
	for _, instruction := range operations {
		result += instruction.Left * instruction.Right
	}
	return result
}
