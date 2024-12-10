package day04_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day04"
	"github.com/stretchr/testify/require"
)

func TestPart1_Small(t *testing.T) {
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

func TestPart1_Example(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	result := day04.Part1(input)

	require.Equal(t, 18, result)
}

func TestPart2_Example(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	result := day04.Part2(input)

	require.Equal(t, 9, result)
}

func TestPart2_MBottom(t *testing.T) {
	input := []string{
		"S.S",
		".A.",
		"M.M",
	}

	result := day04.Part2(input)

	require.Equal(t, 1, result)
}

func TestPart2_MLeft(t *testing.T) {
	input := []string{
		"M.S",
		".A.",
		"M.S",
	}

	result := day04.Part2(input)

	require.Equal(t, 1, result)
}

func TestPart2_MTop(t *testing.T) {
	input := []string{
		"M.M",
		".A.",
		"S.S",
	}

	result := day04.Part2(input)

	require.Equal(t, 1, result)
}

func TestPart2_MRight(t *testing.T) {
	input := []string{
		"S.M",
		".A.",
		"S.M",
	}

	result := day04.Part2(input)

	require.Equal(t, 1, result)
}
