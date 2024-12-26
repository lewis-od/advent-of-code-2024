package day08

import (
	"unicode"
)

type Point struct {
	X, Y int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Subtract(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p Point) IsOnMap(maxX, maxY int) bool {
	return p.X >= 0 && p.X < maxX && p.Y >= 0 && p.Y < maxY
}

func Part1(input []string) int {
	maxX := len(input)
	maxY := len(input[0])

	antennae := FindAntennae(input)
	nodes := findNodes(antennae, maxX, maxY)

	return len(nodes)
}

func findNodes(antennae map[rune][]Point, maxX int, maxY int) map[Point]bool {
	nodes := map[Point]bool{}
	for _, locations := range antennae {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				first := locations[i]
				second := locations[j]
				delta := second.Subtract(first)

				nodeOne := second.Add(delta)
				if nodeOne.IsOnMap(maxX, maxY) {
					nodes[nodeOne] = true
				}

				nodeTwo := first.Subtract(delta)
				if nodeTwo.IsOnMap(maxX, maxY) {
					nodes[nodeTwo] = true
				}
			}
		}
	}
	return nodes
}

func Part2(input []string) int {
	maxX := len(input)
	maxY := len(input[0])

	antennae := FindAntennae(input)
	nodes := findNodesHarmonic(antennae, maxX, maxY)

	return len(nodes)
}

func findNodesHarmonic(antennae map[rune][]Point, maxX int, maxY int) map[Point]bool {
	nodes := map[Point]bool{}
	for _, locations := range antennae {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				first := locations[i]
				second := locations[j]
				delta := second.Subtract(first)

				nodes[first] = true
				nodes[second] = true

				nodeOne := second.Add(delta)
				for nodeOne.IsOnMap(maxX, maxY) {
					nodes[nodeOne] = true
					nodeOne = nodeOne.Add(delta)
				}

				nodeTwo := first.Subtract(delta)
				for nodeTwo.IsOnMap(maxX, maxY) {
					nodes[nodeTwo] = true
					nodeTwo = nodeTwo.Subtract(delta)
				}
			}
		}
	}
	return nodes
}

func FindAntennae(input []string) map[rune][]Point {
	antennae := map[rune][]Point{}
	for y, row := range input {
		for x, col := range row {
			if !unicode.IsLetter(col) && !unicode.IsNumber(col) {
				continue
			}

			point := Point{x, y}
			antennae[col] = append(antennae[col], point)
		}
	}
	return antennae
}
