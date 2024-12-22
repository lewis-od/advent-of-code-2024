package main

import (
	"fmt"
	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day07"
)

func main() {
	input := common.ReadInputFileLines()
	result := day07.Part1(input)
	fmt.Printf("Part 1: %d\n", result)

	result = day07.Part2(input)
	fmt.Printf("Part 2: %d\n", result)
}
