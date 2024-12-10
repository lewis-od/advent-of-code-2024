package day04

import "github.com/lewis-od/aoc24/internal/day04/grid"

func Part1(input []string) int {
	g := grid.NewGrid(input)
	numXmases := 0
	for y := range g.Height() {
		for x := range g.Width() {
			numXmases += g.CountXmasesFrom(x, y)
		}
	}
	return numXmases
}
