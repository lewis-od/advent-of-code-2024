package day08_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day08"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	result := day08.Part1(input)

	require.Equal(t, 14, result)
}

func TestFindAntennae(t *testing.T) {
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	result := day08.FindAntennae(input)

	expected := map[rune][]day08.Point{
		'0': {
			{8, 1},
			{5, 2},
			{7, 3},
			{4, 4},
		},
		'A': {
			{6, 5},
			{8, 8},
			{9, 9},
		},
	}
	require.Equal(t, expected, result)
}
