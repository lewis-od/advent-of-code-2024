package day07_test

import (
	"testing"

	"github.com/lewis-od/aoc24/internal/day07"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	result := day07.Part1(input)

	require.Equal(t, int64(3749), result)
}

func TestPart2(t *testing.T) {
	input := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	result := day07.Part2(input)

	require.Equal(t, int64(11387), result)
}

func TestParseEquation(t *testing.T) {
	input := []string{
		"190: 10 19",
		"3267: 81 40 27",
	}

	equations := day07.ParseEquations(input)

	expectedEquations := []day07.Equation{
		{
			Target:   190,
			Operands: []int64{10, 19},
		},
		{
			Target:   3267,
			Operands: []int64{81, 40, 27},
		},
	}
	require.Equal(t, expectedEquations, equations)
}

func TestEquation_Evaluate(t *testing.T) {
	equation := day07.Equation{
		Operands: []int64{10, 19, 10},
	}
	operations := []day07.Operator{
		day07.Add, day07.Multiply,
	}

	result := equation.Evaluate(operations)

	require.Equal(t, int64(290), result)
}

func TestOperatorPermutations(t *testing.T) {
	permutations := day07.OperatorPermutations(2, []day07.Operator{day07.Add, day07.Multiply})

	expectedPermutations := [][]day07.Operator{
		{day07.Add, day07.Add},
		{day07.Add, day07.Multiply},
		{day07.Multiply, day07.Add},
		{day07.Multiply, day07.Multiply},
	}
	require.Equal(t, expectedPermutations, permutations)
}

func TestEquation_IsSolvable(t *testing.T) {
	equation := day07.Equation{
		Target: 3267,
		Operands: []int64{
			81, 40, 27,
		},
	}
	ops := []day07.Operator{day07.Add, day07.Multiply}

	isSolvable := equation.IsSolvable(ops)

	require.True(t, isSolvable)
}

func TestEquation_IsNotSolvable(t *testing.T) {
	equation := day07.Equation{
		Target: 7290,
		Operands: []int64{
			6, 8, 6, 15,
		},
	}
	ops := []day07.Operator{day07.Add, day07.Multiply}

	isSolvable := equation.IsSolvable(ops)

	require.False(t, isSolvable)
}

func TestEquation_IsSolvableConcat(t *testing.T) {
	equation := day07.Equation{
		Target: 7290,
		Operands: []int64{
			6, 8, 6, 15,
		},
	}
	ops := []day07.Operator{day07.Add, day07.Multiply, day07.Concat}

	isSolvable := equation.IsSolvable(ops)

	require.True(t, isSolvable)
}
