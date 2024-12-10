package main

import (
	"fmt"
	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day04"
)

func main() {
	input := common.ReadInputFileLines()
	numXmases := day04.Part1(input)
	fmt.Printf("Num Xmases: %d\n", numXmases)
}
