package day06

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
		if y >= len(obstacles) {
			break
		}

		x := g.x + dx
		if x >= len(obstacles[0]) {
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

	for {
		guard.Step(obstacles)
		if guard.x >= maxX || guard.y >= maxY {
			break
		}
		visited[guard.y][guard.x] = true
	}

	numVisited := 0
	for _, row := range visited {
		for _, wasVisited := range row {
			if wasVisited {
				numVisited++
			}
		}
	}
	return numVisited
}

func setup(input []string, maxX, maxY int) (visited, obstacles [][]bool, guard *Guard) {
	visited = make([][]bool, maxY)
	obstacles = make([][]bool, maxY)
	for y := range obstacles {
		obstacles[y] = make([]bool, maxX)
		visited[y] = make([]bool, maxX)
		for x := range obstacles[y] {
			point := rune(input[y][x])
			if point == '#' {
				obstacles[y][x] = true
			} else if point == '^' {
				guard = &Guard{direction: North, x: x, y: y}
				visited[y][x] = true
			}
		}
	}
	return
}

//func printState(obstacles, visited [][]bool, guard *Guard) {
//	for y, row := range obstacles {
//		for x, isObstacle := range row {
//			if x == guard.x && y == guard.y {
//				fmt.Print(guard.direction.String())
//			} else if isObstacle {
//				fmt.Print("#")
//			} else if visited[y][x] {
//				fmt.Print("X")
//			} else {
//				fmt.Print(".")
//			}
//		}
//		fmt.Println()
//	}
//	fmt.Println()
//}
