package day05

import (
	"github.com/lewis-od/aoc24/internal/common"
	"slices"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	rules, pagesList := ParseInput(input)

	result := 0
	for _, pages := range pagesList {
		if !ArePagesSorted(rules, pages) {
			continue
		}
		middleValue := pages[len(pages)/2]
		result += middleValue
	}

	return result
}

func Part2(input []string) int {
	rules, pagesList := ParseInput(input)
	compare := generateComparatorFromRules(rules)

	result := 0
	for _, pages := range pagesList {
		if ArePagesSorted(rules, pages) {
			continue
		}
		slices.SortFunc(pages, compare)
		middleValue := pages[len(pages)/2]
		result += middleValue
	}

	return result
}

func ParseInput(input []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	pages := make([][]int, 0)

	parsedAllRules := false
	for _, line := range input {
		if line == "" {
			parsedAllRules = true
			continue
		}

		if !parsedAllRules {
			first, second := parseRule(line)
			rules[first] = append(rules[first], second)
		} else {
			thesePages := parsePages(line)
			pages = append(pages, thesePages)
		}
	}

	return rules, pages
}

func parseRule(line string) (int, int) {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		panic("invalid rule: " + line)
	}

	first := common.Must(strconv.Atoi(parts[0]))
	second := common.Must(strconv.Atoi(parts[1]))
	return first, second
}

func parsePages(line string) []int {
	parts := strings.Split(line, ",")
	pages := make([]int, len(parts))
	for n, part := range parts {
		pages[n] = common.Must(strconv.Atoi(part))
	}
	return pages
}

func ArePagesSorted(rules map[int][]int, pages []int) bool {
	compare := generateComparatorFromRules(rules)
	return slices.IsSortedFunc(pages, compare)
}

func generateComparatorFromRules(rules map[int][]int) func(int, int) int {
	return func(a, b int) int {
		greater := rules[b]
		if slices.Contains(greater, a) {
			return 1
		}
		return -1
	}
}
