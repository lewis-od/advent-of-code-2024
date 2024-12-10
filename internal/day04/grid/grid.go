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
			numXmases += g.countXmasesFromInDirection(x, y, dx, dy)
		}
	}
	return numXmases
}

func (g *Grid) countXmasesFromInDirection(x, y int, dx, dy int) int {
	xmas := "XMAS"
	for n, target := range xmas {
		actual := g.Get(x+n*dx, y+n*dy)
		if actual != target {
			return 0
		}
	}
	return 1
}
