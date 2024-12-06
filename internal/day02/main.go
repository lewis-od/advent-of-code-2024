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

	for i := 0; i < len(report)-1; i++ {
		this := report[i]
		next := report[i+1]
		didIncrease := next > this

		if i == 0 {
			isIncreasing = didIncrease
		}

		difference := common.AbsInt(this - next)
		if difference > 3 || difference < 1 {
			return false
		}

		if isIncreasing != didIncrease {
			return false
		}
	}
	return true
}

func Part2(input []string) int {
	numSafe := 0
	for _, report := range input {
		levels := parseLevels(report)
		isSafe := IsReportSafeWithDampener(levels)
		if isSafe {
			numSafe++
		}
	}

	return numSafe
}

func IsReportSafeWithDampener(report []int) bool {
	if IsReportSafe(report) {
		return true
	}

	possibleReports := combinationsWithElementRemoved(report)
	for _, report := range possibleReports {
		if IsReportSafe(report) {
			return true
		}
	}
	return false
}

func combinationsWithElementRemoved(report []int) [][]int {
	var result [][]int
	for i := 0; i < len(report); i++ {
		newReport := make([]int, 0, len(report)-1)
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		result = append(result, newReport)
	}
	return result
}
