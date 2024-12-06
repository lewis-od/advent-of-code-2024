package day03_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day03"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	result := day03.Part1(input)

	require.Equal(t, 161, result)
}

func TestParseInstructions(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)"

	result := day03.ParseInstructions(input)

	expectedInstructions := []day03.MulInstruction{
		{Left: 2, Right: 4},
		{Left: 5, Right: 5},
	}
	require.Equal(t, expectedInstructions, result)
}
