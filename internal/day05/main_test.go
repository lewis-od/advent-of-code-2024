package day05_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day05"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}

	result := day05.Part1(input)

	require.Equal(t, 143, result)
}

func TestParseInput(t *testing.T) {
	input := []string{
		"75|13",
		"53|13",
		"53|20",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
	}

	rules, pages := day05.ParseInput(input)

	expectedRules := map[int][]int{
		75: {13},
		53: {13, 20},
	}
	expectedPages := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
	}
	require.Equal(t, rules, expectedRules)
	require.Equal(t, pages, expectedPages)
}

func TestArePagesSorted_Sorted(t *testing.T) {
	rules := map[int][]int{
		47: {53, 13, 61, 29},
		97: {13, 61, 47, 29, 53, 75},
		75: {29, 53, 47, 61, 13},
		61: {13, 53, 29},
		29: {13},
		53: {29, 13},
	}
	pages := []int{75, 47, 61, 53, 29}

	areSorted := day05.ArePagesSorted(rules, pages)

	require.True(t, areSorted)
}

func TestArePagesSorted_NotSorted(t *testing.T) {
	rules := map[int][]int{
		47: {53, 13, 61, 29},
		97: {13, 61, 47, 29, 53, 75},
		75: {29, 53, 47, 61, 13},
		61: {13, 53, 29},
		29: {13},
		53: {29, 13},
	}
	pages := []int{75, 97, 47, 61, 53}

	areSorted := day05.ArePagesSorted(rules, pages)

	require.False(t, areSorted)
}

func TestPart2(t *testing.T) {
	input := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}

	result := day05.Part2(input)

	require.Equal(t, 123, result)
}
