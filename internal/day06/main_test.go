package day06_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day06"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	result := day06.Part1(input)

	require.Equal(t, 41, result)
}

func TestPart2(t *testing.T) {
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	result := day06.Part2(input)

	require.Equal(t, 6, result)
}
