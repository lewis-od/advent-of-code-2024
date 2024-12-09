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
		if instruction.Type != parsing.OpMul {
			continue
		}
		result += instruction.Left * instruction.Right
	}
	return result
}

func Part2(input string) int {
	scanner := parsing.NewScanner(input)
	parser := parsing.NewParser(scanner.ScanTokens())
	operations := parser.Parse()

	result := 0
	isMulEnabled := true
	for _, instruction := range operations {
		// Gross
		if instruction.Type == parsing.OpDo {
			isMulEnabled = true
		} else if instruction.Type == parsing.OpDont {
			isMulEnabled = false
		} else if instruction.Type == parsing.OpMul && isMulEnabled {
			result += instruction.Left * instruction.Right
		}
	}
	return result
}
