package util

import (
	"fmt"
)

// ----------------------
// Point
// ----------------------

type GridPoint struct {
	x       int
	y       int
	value   string
	visited bool
}

func NewPoint(x, y int) GridPoint {
	p := GridPoint{}
	p.x = x
	p.y = y
	return p
}

func NewPointV(x, y int, v string) GridPoint {
	p := GridPoint{}
	p.x = x
	p.y = y
	p.value = v
	return p
}

func (self GridPoint) WithValue(v string) GridPoint {
	self.value = v
	return self
}

func (self GridPoint) LocationEqual(p GridPoint) bool {
	return self.x == p.x && self.y == p.y
}

func (self GridPoint) Equal(p GridPoint) bool {
	return self.x == p.x && self.y == p.y && self.value == p.value
}

func (self GridPoint) ValueEqual(p GridPoint) bool {
	return self.value == p.value
}

func (self GridPoint) Sub(p GridPoint) GridPoint {
	return GridPoint{self.x - p.x, self.y - p.y, "", false}
}

func (self GridPoint) IsIn(list GridRow) bool {
	for _, p := range list {
		if self == p {
			return true
		}
	}
	return false
}

// ----------------------
// Row
// ----------------------

type GridRow []GridPoint

func (e GridRow) Len() int {
	return len(e)
}

func (e GridRow) Less(i, j int) bool {
	if e[i].y < e[j].y {
		return true
	}
	if e[i].y > e[j].y {
		return false
	}
	return e[i].x < e[j].x
}

func (e GridRow) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (self *GridRow) Contains(p GridPoint) bool {
	for _, x := range *self {
		if x.LocationEqual(p) {
			return true
		}
	}
	return false
}

func (self *GridRow) ContainsValue(p GridPoint) bool {
	for _, x := range *self {
		if x.ValueEqual(p) {
			return true
		}
	}
	return false
}

// ----------------------
// Grid
// ----------------------

type Grid struct {
	grid    []string
	current GridPoint
	max_row int
	max_col int
}

func MakeGrid(s []string) Grid {
	g := Grid{
		s,
		NewPoint(0, 0),
		len(s) - 1,
		len(s[0]) - 1,
	}
	return g
}

func (g *Grid) MaxRow() int {
	return (*g).max_row
}

func (g *Grid) MaxCol() int {
	return (*g).max_col
}

func (g *Grid) Get(x, y int) GridPoint {
	value := (*g).grid[y][x]
	return GridPoint{x, y, string(value), false}
}

func (g *Grid) GetNeighbours(p GridPoint, checkCorners bool) (ans GridRow) {
	for _, roff := range []int{-1, 0, 1} {
		x := p.x + roff
		if x < 0 || x >= g.MaxRow() {
			// fmt.Printf("side %d\n", x)
			continue
		}

		for _, coff := range []int{-1, 0, 1} {
			y := p.y + coff
			if y < 0 || y >= g.MaxCol() {
				// fmt.Printf("side %d,%d\n", x, y)
				continue
			}

			candidate := g.Get(x, y)
			if candidate.LocationEqual(p) {
				continue
			}
			// fmt.Printf(%v\n", candidate)

			// Dont check corners and same point
			corners := GridRow{NewPoint(0, 0), NewPoint(-1, -1), NewPoint(1, -1), NewPoint(-1, 1), NewPoint(1, 1)}
			if !checkCorners && p.Sub(candidate).IsIn(corners) {
				// fmt.Printf("skip corner %d,%d\n", x, y)
				continue
			}

			// fmt.Printf("%d,%d\n", x, y)
			ans = append(ans, candidate)
			// fmt.Printf("%v", ans)
		}
	}
	return ans
}

func (self Grid) Print() {
	// Print final grid
	for _, s := range self.grid {
		fmt.Printf("%s\n", s)
	}
}

func (self Grid) FindValue(value string) (GridPoint, error) {
	for y := range self.grid {
		for x := range self.grid[y] {
			if string(self.grid[y][x]) == value {
				return NewPoint(x, y), nil
			}
		}
	}
	return GridPoint{}, fmt.Errorf("%s : value not found.", value)
}

func (self Grid) FindValues(value string) GridRow {
	r := GridRow{}
	for y := range self.grid {
		for x := range self.grid[y] {
			if string(self.grid[y][x]) == value {
				r = append(r, NewPoint(x, y))
			}
		}
	}
	return r
}
