package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	// "github.com/louisdcoulombe/advent-of-code-go/cast"
	"github.com/louisdcoulombe/advent-of-code-go/util"
)

//go:embed input.txt
var input string

const (
	RED   string = "red"
	GREEN string = "green"
	BLUE  string = "blue"
)

const (
	MAX_RED   int = 12
	MAX_GREEN int = 13
	MAX_BLUE  int = 14
)

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
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

func part1(input string) int {
	parsed := parseInput(input)
	validGames := []int{}
GameLoop:
	for gameIdx, game := range parsed {
		// parse game
		parts := strings.Split(game, ":")
		// parse variable num turns
		for _, turn := range strings.Split(parts[1], ";") {
			for _, item := range strings.Split(turn, ", ") {
				parts = strings.Split(strings.Trim(item, " "), " ")
				value, _ := strconv.Atoi(parts[0])
				// fmt.Printf("%s = %d\n", item, value)
				switch parts[1] {
				case RED:
					if value > MAX_RED {
						continue GameLoop
					}
				case GREEN:
					if value > MAX_GREEN {
						continue GameLoop
					}
				case BLUE:
					if value > MAX_BLUE {
						continue GameLoop
					}
				}
			}
		}

		validGames = append(validGames, gameIdx+1)

	}

	// fmt.Printf("%d", validGames)
	sum := 0
	for _, v := range validGames {
		sum += v
	}
	return sum
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
