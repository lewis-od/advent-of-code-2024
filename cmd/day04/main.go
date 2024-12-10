package main

import (
	"fmt"
	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day04"
)

func main() {
	input := common.ReadInputFileLines()
	numXmases := day04.Part1(input)
	fmt.Printf("Num XMASes: %d\n", numXmases)

	numCrossMases := day04.Part2(input)
	fmt.Printf("Num X-MASes: %d\n", numCrossMases)
}
