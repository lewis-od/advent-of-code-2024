package main

import (
	"fmt"
	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day05"
)

func main() {
	input := common.ReadInputFileLines()
	result := day05.Part1(input)
	fmt.Printf("Part 1: %d\n", result)
}
