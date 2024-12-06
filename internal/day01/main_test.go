package day01_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day01"
	"github.com/stretchr/testify/require"
)

func TestParseList(t *testing.T) {
	input := "3   4\n4   3\n35039   67568\n125   9813"

	left, right := day01.ParseLists(string(input))

	expectedLeft := []int{3, 4, 35039, 125}
	expectedRight := []int{4, 3, 67568, 9813}
	require.Equal(t, expectedLeft, left)
	require.Equal(t, expectedRight, right)
}

func TestCalculateDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 9, 3, 5, 3}

	distance := day01.CalculateDistance(left, right)

	require.Equal(t, 11, distance)
}

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	distance := day01.Part1(input)

	require.Equal(t, 11, distance)
}
