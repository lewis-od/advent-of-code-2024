package day03

import (
	"regexp"
	"strconv"

	"github.com/lewis-od/aoc24/internal/common"
)

var mulRegex = regexp.MustCompile(`mul\((?P<left>\d+),(?P<right>\d+)\)`)

type MulInstruction struct {
	Left  int
	Right int
}

func (instruction MulInstruction) Execute() int {
	return instruction.Left * instruction.Right
}

func Part1(input string) int {
	instructions := ParseInstructions(input)

	result := 0
	for _, instruction := range instructions {
		result += instruction.Execute()
	}
	return result
}

func ParseInstructions(input string) []MulInstruction {
	matches := mulRegex.FindAllStringSubmatch(input, -1)

	var instructions []MulInstruction
	for _, match := range matches {
		instruction := MulInstruction{}
		for i, groupName := range mulRegex.SubexpNames() {
			if i == 0 {
				continue // Skip first element (full match)
			}

			if groupName == "left" {
				instruction.Left = common.Must(strconv.Atoi(match[i]))
			} else if groupName == "right" {
				instruction.Right = common.Must(strconv.Atoi(match[i]))
			}
		}
		instructions = append(instructions, instruction)
	}
	return instructions
}
