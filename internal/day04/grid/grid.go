package grid

type Grid struct {
	values []string
}

func NewGrid(values []string) *Grid {
	return &Grid{values}
}

func (g *Grid) Get(x, y int) rune {
	if y < 0 || y >= len(g.values) {
		return '.'
	}
	row := g.values[y]
	if x < 0 || x >= len(row) {
		return '.'
	}
	return rune(row[x])
}

func (g *Grid) Width() int {
	row := g.values[0]
	return len(row)
}

func (g *Grid) Height() int {
	return len(g.values)
}

func (g *Grid) CountXmasesFrom(x, y int) int {
	numXmases := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			if g.isWordInDirection("XMAS", x, y, dx, dy) {
				numXmases++
			}
		}
	}
	return numXmases
}

func (g *Grid) CountCrossMasesFrom(x, y int) int {
	numCrossMases := 0
	if g.isMLeftCrossMas(x, y) {
		numCrossMases++
	}

	if g.isMRightCrossMas(x, y) {
		numCrossMases++
	}

	if g.isMBottomCrossMas(x, y) {
		numCrossMases++
	}

	if g.isMTopCrossMas(x, y) {
		numCrossMases++
	}
	return numCrossMases
}

func (g *Grid) isMLeftCrossMas(x, y int) bool {
	if !g.isWordInDirection("MAS", x, y, 1, 1) {
		return false
	}
	return g.isWordInDirection("MAS", x, y+2, 1, -1)
}

func (g *Grid) isMRightCrossMas(x, y int) bool {
	if !g.isWordInDirection("MAS", x, y, -1, 1) {
		return false
	}
	return g.isWordInDirection("MAS", x, y+2, -1, -1)
}

func (g *Grid) isMBottomCrossMas(x, y int) bool {
	if !g.isWordInDirection("MAS", x, y, -1, -1) {
		return false
	}
	return g.isWordInDirection("MAS", x-2, y, 1, -1)
}

func (g *Grid) isMTopCrossMas(x, y int) bool {
	if !g.isWordInDirection("MAS", x, y, 1, 1) {
		return false
	}
	return g.isWordInDirection("MAS", x+2, y, -1, 1)
}

func (g *Grid) isWordInDirection(word string, x, y int, dx, dy int) bool {
	for n, target := range word {
		actual := g.Get(x+n*dx, y+n*dy)
		if actual != target {
			return false
		}
	}
	return true
}
