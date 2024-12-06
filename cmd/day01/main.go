package main

import (
	"fmt"

	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day01"
)

func main() {
	input := common.ReadInputFile()
	distance := day01.Part1(string(input))
	fmt.Printf("The distance is %d\n", distance)
}
