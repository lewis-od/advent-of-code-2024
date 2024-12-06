package main

import (
	"fmt"

	"github.com/lewis-od/aoc24/internal/common"
	"github.com/lewis-od/aoc24/internal/day02"
)

func main() {
	reports := common.ReadInputFileLines()
	numSafe := day02.Part1(reports)
	fmt.Printf("Safe reports: %d\n", numSafe)
}
