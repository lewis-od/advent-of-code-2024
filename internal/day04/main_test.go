package day04_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day04"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := []string{
		"..X...",
		".SAMX.",
		".A..A.",
		"XMAS.S",
		".X....",
	}

	result := day04.Part1(input)

	require.Equal(t, 4, result)
}
