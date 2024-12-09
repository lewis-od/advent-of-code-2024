package main

import (
	"fmt"

	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day03"
)

func main() {
	memory := common.ReadInputFile()
	result := day03.Part1(string(memory))
	fmt.Printf("Result 1: %d\n", result)

	result2 := day03.Part2(string(memory))
	fmt.Printf("Result 2: %d\n", result2)
}
