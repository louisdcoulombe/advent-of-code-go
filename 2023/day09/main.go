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

type History []int

func parseHistories(in []string) (ans []History) {
	for _, l := range in {
		ans = append(ans, util.StringsToInts(l, " "))
	}
	return ans
}

func calculateDiffs(h History) (ans History) {
	for i := 1; i < len(h); i++ {
		ans = append(ans, h[i]-h[i-1])
	}
	return ans
}

func extrapolate(h History) int {
	fmt.Printf("%v\n", h)
	if util.Sum[int, int](h) == 0 {
		return 0
	}
	return h[len(h)-1] + extrapolate(calculateDiffs(h))
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	histories := parseHistories(parsed)
	fmt.Printf("%v", histories)

	ans := []int{}
	for _, h := range histories {
		x := extrapolate(h)
		fmt.Printf("=%d\n", x)
		ans = append(ans, x)

	}
	return util.Sum[int, int](ans)
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
