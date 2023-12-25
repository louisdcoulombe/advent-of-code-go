package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/louisdcoulombe/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

type (
	PointList []Point
	Point     struct {
		x int
		y int
	}
)

func (e PointList) Len() int {
	return len(e)
}

func (e PointList) Less(i, j int) bool {
	if e[i].y < e[j].y {
		return true
	}
	if e[i].y > e[j].y {
		return false
	}
	return e[i].x < e[j].x
}

func (e PointList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (self Point) Sub(p Point) Point {
	return Point{self.x - p.x, self.y - p.y}
}

func (self Point) IsIn(list []Point) bool {
	for _, p := range list {
		if self == p {
			return true
		}
	}
	return false
}

func makeGrid(s []string) Grid {
	g := Grid{
		s,
		Point{0, 0},
		len(s) - 1,
		len(s[0]) - 1,
	}
	return g
}

type Grid struct {
	grid    []string
	current Point
	max_row int
	max_col int
}

func (self Grid) Print() {
	// Print final grid
	for _, s := range self.grid {
		fmt.Printf("%s\n", s)
	}
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (g *Grid) checkAround(p Point) (ans []Point) {
	corners := []Point{{0, 0}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
	for _, roff := range []int{-1, 0, 1} {
		x := p.x + roff
		if x < 0 || x > g.max_row {
			continue
		}

		for _, coff := range []int{-1, 0, 1} {
			y := p.y + coff
			if y < 0 || y > g.max_col {
				continue
			}

			candidate := Point{x, y}

			// Dont check corners and same point
			if p.Sub(candidate).IsIn(corners) {
				continue
			}

			candidate_symbol := string(g.grid[y][x])
			current := string(g.grid[p.y][p.x])
			diff := candidate.Sub(p)
			fmt.Printf(" '%s'%v - '%s'%v = %v", candidate_symbol, candidate, current, p, diff)

			indexes, ok := SYMBOL_GO[candidate_symbol]
			if !ok {
				fmt.Printf("\n")
				continue
			}

			if diff.IsIn(indexes) {
				fmt.Printf(" ADDED")
				ans = append(ans, candidate)
			}
			fmt.Printf("\n")
		}
	}
	return ans
}

// Where you are allowed to come from (Symbol position - Current position)
//
//	N (0, 1)
//	 W (1, 0)
//	 E (-1, 0)
//	S (0, -1)
var SYMBOL_GO = map[string][]Point{
	"|": {{0, -1}, {0, 1}},  // is a vertical pipe connecting north (0, -1) and south(0, 1)
	"-": {{-1, 0}, {1, 0}},  // is a horizontal pipe connecting east (-1,0) and west (1,0)
	"L": {{0, 1}, {-1, 0}},  // is a 90-degree bend connecting north(0, -1)  and east (-1, 0)
	"J": {{0, 1}, {1, 0}},   // is a 90-degree bend connecting north(0, -1)  and west (1,0 )
	"7": {{0, -1}, {1, 0}},  // is a 90-degree bend connecting south(0, 1)  and wes (1, 0)
	"F": {{0, -1}, {-1, 0}}, // is a 90-degree bend connecting south(0, 1)  and east (-1, 0)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		_ = util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		_ = util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func findStartIndex(g []string) Point {
	for y := range g {
		for x := range g[y] {
			if string(g[y][x]) == "S" {
				return Point{x, y}
			}
		}
	}
	panic("Start not found")
}

type Queue []Point

func (q *Queue) Push(p Point) {
	*q = append(*q, p)
}

func (q *Queue) Pop() (val Point) {
	if len(*q) == 0 {
		panic("Empty queue")
	}

	val = (*q)[0]
	*q = (*q)[1:]
	return val
}

func part1(input string) int {
	in := parseInput(input)
	return GetTilesEnclosed(in)
}

func GetTilesEnclosed(input []string) int {
	g := makeGrid(input)

	start := findStartIndex(g.grid)
	fmt.Printf("\nS=(%d,%d)\n", start.x, start.y)

	queue := Queue{}
	for _, pts := range g.checkAround(start) {
		queue = append(queue, pts)
	}

	g.current = start
	counter := 0
	putNumber := func(p Point) {
		str := fmt.Sprint(counter % 9)
		if str == "7" {
			str = "X"
		}
		g.grid[p.y] = util.ReplaceAtIndex(g.grid[p.y], str, p.x)
	}

	putNumber(start)
	if len(queue) != 2 {
		panic("Not 2 items")
	}

	for len(queue) > 0 {
		counter++
		// Check first side
		current := queue.Pop()
		fmt.Printf("-- %v\n", current)
		values := g.checkAround(current)
		if len(values) == 0 && len(queue) == 0 {
			fmt.Printf("Last %d", counter)
			break
		}
		if len(values) == 0 {
			fmt.Printf("Len=0\n")
			continue
		}
		// Append next value
		queue.Push(values[0])
		putNumber(current)
		g.Print()

		// Pop the other side
		current = queue.Pop()
		values = g.checkAround(current)
		if len(values) == 0 && len(queue) == 0 {
			fmt.Printf("Last 2 = %d", counter)
			break
		}
		if len(values) == 0 {
			fmt.Printf("Len=0\n")
			continue
		}

		// Append next value
		queue.Push(values[0])
		putNumber(current)
		g.Print()
	}

	return counter - 1
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
