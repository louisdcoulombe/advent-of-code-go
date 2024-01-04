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

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func (g *Grid) checkAround(p Point) (ans []Point) {
// 	corners := []Point{{0, 0}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
// 	for _, roff := range []int{-1, 0, 1} {
// 		x := p.x + roff
// 		if x < 0 || x > g.max_row {
// 			continue
// 		}
//
// 		for _, coff := range []int{-1, 0, 1} {
// 			y := p.y + coff
// 			if y < 0 || y > g.max_col {
// 				continue
// 			}
//
// 			candidate := Point{x, y}
//
// 			// Dont check corners and same point
// 			if p.Sub(candidate).IsIn(corners) {
// 				continue
// 			}
//
// 			candidate_symbol := string(g.grid[y][x])
// 			diff := candidate.Sub(p)
// 			// current := string(g.grid[p.y][p.x])
// 			// fmt.Printf(" '%s'%v - '%s'%v = %v", candidate_symbol, candidate, current, p, diff)
//
// 			indexes, ok := SYMBOL_GO[candidate_symbol]
// 			if !ok {
// 				// fmt.Printf("\n")
// 				continue
// 			}
//
// 			if diff.IsIn(indexes) {
// 				// fmt.Printf(" ADDED")
// 				ans = append(ans, candidate)
// 			}
// 			// fmt.Printf("\n")
// 		}
// 	}
// 	return ans
// }

// Where you are allowed to come from (Symbol position - Current position)
//
//	N (0, 1)
//	 W (1, 0)
//	 E (-1, 0)
//	S (0, -1)
var SYMBOL_GO = map[string]util.GridRow{
	"|": {util.NewPoint(0, -1), util.NewPoint(0, 1)},  // is a vertical pipe connecting north (0, -1) and south(0, 1)
	"-": {util.NewPoint(-1, 0), util.NewPoint(1, 0)},  // is a horizontal pipe connecting east (-1,0) and west (1,0)
	"L": {util.NewPoint(0, 1), util.NewPoint(-1, 0)},  // is a 90-degree bend connecting north(0, -1)  and east (-1, 0)
	"J": {util.NewPoint(0, 1), util.NewPoint(1, 0)},   // is a 90-degree bend connecting north(0, -1)  and west (1,0 )
	"7": {util.NewPoint(0, -1), util.NewPoint(1, 0)},  // is a 90-degree bend connecting south(0, 1)  and wes (1, 0)
	"F": {util.NewPoint(0, -1), util.NewPoint(-1, 0)}, // is a 90-degree bend connecting south(0, 1)  and east (-1, 0)
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

type Queue []util.GridPoint

func (q *Queue) Push(p util.GridPoint) {
	*q = append(*q, p)
}

func (q *Queue) Pop() (val util.GridPoint) {
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

func filterSymbols(current util.GridPoint, list util.GridRow) (ans util.GridRow) {
	for _, candidate := range list {
		diff := candidate.Sub(current)
		indexes, ok := SYMBOL_GO[candidate.Value()]
		if !ok {
			continue
		}
		if diff.IsIn(indexes) {
			ans = append(ans, candidate)
		}
	}
	return ans
}

func GetTilesEnclosed(input []string) int {
	g := util.MakeGrid(input)
	start, err := g.FindValue("S")
	if err != nil {
		panic("Start not found")
	}
	queue := Queue{}
	path := util.GridRow{}
	queue.Push(start)

	for len(queue) > 0 {
		current := queue.Pop()
		if current.IsIn(path) {
			break
		}

		path = append(path, current)
		values := g.GetNeighbours(current, false)
		values = filterSymbols(current, values)
		for _, v := range values {
			queue.Push(v)
		}
	}

	fmt.Printf("%v\n", path)
	return len(path)
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
