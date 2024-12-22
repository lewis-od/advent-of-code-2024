package day07

import (
	"fmt"
	"github.com/lewis-od/aoc24/internal/common"
	"strconv"
	"strings"
)

type Operator uint8

const (
	Add Operator = iota
	Multiply
	Concat
)

type Equation struct {
	Target   int64
	Operands []int64
}

func (e Equation) IsSolvable(operators []Operator) bool {
	permutations := OperatorPermutations(len(e.Operands)-1, operators)
	for _, operators := range permutations {
		result := e.Evaluate(operators)
		if result == e.Target {
			return true
		}
	}

	return false
}

func (e Equation) Evaluate(operations []Operator) int64 {
	if len(operations) != len(e.Operands)-1 {
		panic("invalid operation length")
	}

	total := e.Operands[0]
	for n, operation := range operations {
		next := e.Operands[n+1]
		switch operation {
		case Add:
			total += next
		case Multiply:
			total *= next
		case Concat:
			left := fmt.Sprintf("%d", total)
			newTotal := string(strconv.AppendInt([]byte(left), next, 10))
			total = common.Must(strconv.ParseInt(newTotal, 10, 64))
		}
	}

	return total
}

func Part1(input []string) int64 {
	equations := ParseEquations(input)
	operators := []Operator{Add, Multiply}

	result := int64(0)
	for _, equation := range equations {
		if !equation.IsSolvable(operators) {
			continue
		}
		result += equation.Target
	}

	return result
}

func Part2(input []string) int64 {
	equations := ParseEquations(input)
	operators := []Operator{Add, Multiply, Concat}

	result := int64(0)
	for _, equation := range equations {
		if !equation.IsSolvable(operators) {
			continue
		}
		result += equation.Target
	}

	return result
}

func ParseEquations(input []string) []Equation {
	equations := make([]Equation, len(input))
	for n, line := range input {
		equations[n] = parseEquation(line)
	}
	return equations
}

func parseEquation(input string) Equation {
	parts := strings.Split(input, ": ")
	if len(parts) != 2 {
		panic("invalid input: " + input)
	}

	target := common.Must(strconv.ParseInt(parts[0], 10, 64))

	operandStrs := strings.Split(parts[1], " ")
	operands := make([]int64, len(operandStrs))
	for n, str := range operandStrs {
		operands[n] = common.Must(strconv.ParseInt(str, 10, 64))
	}

	return Equation{
		Target:   target,
		Operands: operands,
	}
}

func OperatorPermutations(n int, allOperators []Operator) [][]Operator {
	var result [][]Operator

	var permute func(comb []Operator)
	permute = func(comb []Operator) {
		if len(comb) == n {
			combCopy := append([]Operator{}, comb...)
			result = append(result, combCopy)
			return
		}

		for _, item := range allOperators {
			comb = append(comb, item)
			permute(comb)
			comb = comb[:len(comb)-1]
		}
	}

	permute([]Operator{})
	return result
}
