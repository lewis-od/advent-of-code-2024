package day01

import (
	"slices"
	"strconv"
	"strings"

	"github.com/lewis-od/aoc24/internal/common"
)

func Part1(input string) int {
	left, right := ParseLists(input)
	return CalculateDistance(left, right)
}

func ParseLists(input string) (left []int, right []int) {
	lines := strings.Split(string(input), "\n")
	left = make([]int, len(lines))
	right = make([]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			panic("Expected 2 numbers on each row")
		}
		leftNum := common.Must(strconv.Atoi(parts[0]))
		rightNum := common.Must(strconv.Atoi(parts[1]))

		left[i] = leftNum
		right[i] = rightNum
	}
	return
}

func CalculateDistance(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0
	for i, leftNum := range left {
		rightNum := right[i]
		distance := abs(leftNum - rightNum)
		totalDistance += distance
	}

	return totalDistance
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
