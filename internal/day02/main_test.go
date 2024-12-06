package day02_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day02"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	result := day02.Part1(input)

	require.Equal(t, 2, result)
}

func TestIsReportSafe_Safe(t *testing.T) {
	report := []int{7, 6, 4, 2, 1}

	isSafe := day02.IsReportSafe(report)

	require.True(t, isSafe)
}

func TestIsReportSafe_IncreaseTooLarge(t *testing.T) {
	report := []int{1, 2, 7, 8, 9}

	isSafe := day02.IsReportSafe(report)

	require.False(t, isSafe)
}

func TestIsReportSafe_DecreaseTooLarge(t *testing.T) {
	report := []int{9, 7, 6, 2, 1}

	isSafe := day02.IsReportSafe(report)

	require.False(t, isSafe)
}

func TestIsReportSafe_IncreaseThenDecrease(t *testing.T) {
	report := []int{1, 3, 2, 4, 5}

	isSafe := day02.IsReportSafe(report)

	require.False(t, isSafe)
}

func TestIsReportSafe_RepeatedValue(t *testing.T) {
	report := []int{8, 6, 4, 4, 1}

	isSafe := day02.IsReportSafe(report)

	require.False(t, isSafe)
}

func TestPart2(t *testing.T) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	result := day02.Part2(input)

	require.Equal(t, 4, result)
}

func TestIsReportSafeWithDampener_RepeatedValue(t *testing.T) {
	report := []int{8, 6, 4, 4, 1}

	isSafe := day02.IsReportSafeWithDampener(report)

	require.True(t, isSafe)
}

func TestIsReportSafeWithDampener_IncreaseThenDecrease(t *testing.T) {
	report := []int{1, 3, 2, 4, 5}

	isSafe := day02.IsReportSafeWithDampener(report)

	require.True(t, isSafe)
}
