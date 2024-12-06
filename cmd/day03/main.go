package main

import (
	"fmt"

	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day03"
)

func main() {
	memory := common.ReadInputFile()
	result := day03.Part1(string(memory))
	fmt.Printf("Result: %d\n", result)
}
