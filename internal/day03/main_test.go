package day03_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day03"
	"github.com/lewis-od/aoc24/internal/day03/parsing"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	result := day03.Part1(input)

	require.Equal(t, 161, result)
}

func TestParserAndScanner(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)"

	scanner := parsing.NewScanner(input)
	tokens := scanner.ScanTokens()
	parser := parsing.NewParser(tokens)
	ops := parser.Parse()

	expectedOps := []parsing.Operation{
		{Type: parsing.OpMul, Left: 2, Right: 4},
		{Type: parsing.OpMul, Left: 5, Right: 5},
	}
	require.Equal(t, expectedOps, ops)
}

func TestPart2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	result := day03.Part2(input)

	require.Equal(t, 48, result)
}
