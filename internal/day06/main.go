package day06

import "slices"

type Direction uint8

const (
	North Direction = iota
	East            = iota
	South           = iota
	West            = iota
)

func (d Direction) RotateClockwise() Direction {
	return (d + 1) % 4
}

func (d Direction) Step() (int, int) {
	dx := 0
	dy := 0
	switch d {
	case North:
		dy = -1
	case East:
		dx = 1
	case South:
		dy = 1
	case West:
		dx = -1
	}
	return dx, dy
}

func (d Direction) String() string {
	switch d {
	case North:
		return "^"
	case East:
		return ">"
	case South:
		return "v"
	case West:
		return "<"
	}
	panic("invalid direction")
}

type Guard struct {
	direction Direction
	x, y      int
}

func (g *Guard) Step(obstacles [][]bool) {
	dx, dy := g.direction.Step()

	for {
		y := g.y + dy
		if y >= len(obstacles) || y < 0 {
			break
		}

		x := g.x + dx
		if x >= len(obstacles[0]) || x < 0 {
			break
		}

		if obstacles[g.y+dy][g.x+dx] {
			g.direction = g.direction.RotateClockwise()
			dx, dy = g.direction.Step()
		} else {
			break
		}
	}

	g.x = g.x + dx
	g.y = g.y + dy
}

func Part1(input []string) int {
	maxY := len(input)
	maxX := len(input[0])

	visited, obstacles, guard := setup(input, maxX, maxY)
	if guard == nil {
		panic("no guard found")
	}

	tracePath(guard, obstacles, maxX, maxY, visited)

	numVisited := 0
	for _, row := range visited {
		for _, visitedDirections := range row {
			if len(visitedDirections) > 0 {
				numVisited++
			}
		}
	}
	return numVisited
}

func Part2(input []string) int {
	// Trace path
	// For each point on the path
	//   Place an obstacle at this point
	//   Run the simulation again
	//   Break when:
	//     Leaving map => no loop
	//     Re-visiting already visited square whist facing in same direction as previously => loop
	maxY := len(input)
	maxX := len(input[0])

	visited, obstacles, guard := setup(input, maxX, maxY)
	if guard == nil {
		panic("no guard found")
	}

	startX, startY := guard.x, guard.y
	tracePath(guard, obstacles, maxX, maxY, visited)

	numLoops := 0
	for y, row := range visited {
		for x, visitedDirections := range row {
			if len(visitedDirections) == 0 {
				continue
			}

			obstacles[y][x] = true
			if doesGuardLoop(obstacles, startX, startY, maxX, maxY) {
				numLoops++
			}
			obstacles[y][x] = false
		}
	}

	return numLoops
}

func setup(input []string, maxX, maxY int) (visited [][][]Direction, obstacles [][]bool, guard *Guard) {
	visited = make([][][]Direction, maxY)
	obstacles = make([][]bool, maxY)
	for y := range obstacles {
		obstacles[y] = make([]bool, maxX)
		visited[y] = make([][]Direction, maxX)
		for x := range obstacles[y] {
			point := rune(input[y][x])
			if point == '#' {
				obstacles[y][x] = true
			} else if point == '^' {
				guard = &Guard{direction: North, x: x, y: y}
				visited[y][x] = []Direction{North}
			}
		}
	}
	return
}

func doesGuardLoop(obstacles [][]bool, x, y, maxX, maxY int) bool {
	visited := make([][][]Direction, maxY)
	for y := range obstacles {
		visited[y] = make([][]Direction, maxX)
	}

	guard := &Guard{direction: North, x: x, y: y}
	for {
		guard.Step(obstacles)

		if guard.x >= maxX || guard.x < 0 || guard.y >= maxY || guard.y < 0 {
			return false
		}

		visitedDirections := visited[guard.y][guard.x]
		if slices.Contains(visitedDirections, guard.direction) {
			return true
		}

		visited[guard.y][guard.x] = append(visited[guard.y][guard.x], guard.direction)
	}
}

func tracePath(guard *Guard, obstacles [][]bool, maxX int, maxY int, visited [][][]Direction) {
	for {
		guard.Step(obstacles)
		if guard.x >= maxX || guard.y >= maxY {
			break
		}
		visited[guard.y][guard.x] = append(visited[guard.y][guard.x], guard.direction)
	}
}
