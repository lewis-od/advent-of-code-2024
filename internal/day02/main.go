package day02

import (
	"strconv"
	"strings"

	"github.com/lewis-od/aoc24/internal/common"
)

func Part1(input []string) int {

	numSafe := 0
	for _, report := range input {
		levels := parseLevels(report)
		isSafe := IsReportSafe(levels)
		if isSafe {
			numSafe++
		}
	}

	return numSafe
}

func parseLevels(row string) []int {
	levelStrings := strings.Split(row, " ")
	levels := make([]int, len(levelStrings))
	for i, value := range levelStrings {
		levels[i] = common.Must(strconv.Atoi(value))
	}
	return levels
}

func IsReportSafe(report []int) bool {
	isIncreasing := false
	for i, level := range report {
		if i == 0 {
			continue
		}

		prev := report[i-1]
		diff := common.AbsInt(prev - level)
		if diff > 3 || diff < 1 {
			return false
		}

		if i == 1 {
			isIncreasing = level > prev
		} else {
			didIncrease := level > prev
			if isIncreasing != didIncrease {
				return false
			}
		}
	}
	return true
}
