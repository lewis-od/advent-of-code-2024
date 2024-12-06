package main

import (
	"fmt"

	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day01"
)

func main() {
	input := common.ReadInputFileLines()
	distance := day01.Part1(input)
	fmt.Printf("The distance is %d\n", distance)

	similarity := day01.Part2(input)
	fmt.Printf("The similarity is %d\n", similarity)
}
