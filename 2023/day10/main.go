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

//	N
//
// W E
//
//	S
const (
	VER     = "|" // is a vertical pipe connecting north and south.
	HOR     = "-" // is a horizontal pipe connecting east and west.
	T_RIGHT = "L" // is a 90-degree bend connecting north and east.
	T_LEFT  = "J" // is a 90-degree bend connecting north and west.
	B_LEFT  = "7" // is a 90-degree bend connecting south and west.
	B_RIGHT = "F" // is a 90-degree bend connecting south and east.
	GROUND  = "." // is ground; there is no pipe in this tile.
	START   = "S" // is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
)

var DIRECTIONS = map[string][]int{
	// x, y
	VER:     {0, 1},
	HOR:     {1, 0},
	T_RIGHT: {1, -1},
	T_LEFT:  {-1, -1},
	B_LEFT:  {-1, 1},
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

func findStartIndex(g []string) (int, int) {
	for y := range g {
		for x, c := range g[y] {
			fmt.Printf("%s", string(c))
			if string(g[y][x]) == "S" {
				return x, y
			}
		}
		fmt.Printf("\n")
	}
	// panic("Start not found")
	return 0, 0
}

func part1(input string) int {
	grid := parseInput(input)
	x, y := findStartIndex(grid)
	fmt.Printf("\nS=(%d,%d)\n", x, y)

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
