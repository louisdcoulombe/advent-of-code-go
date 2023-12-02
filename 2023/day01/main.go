package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"

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

func part1(input string) int {
	parsed := parseInput(input)
	sum := 0
	for _, line := range parsed {
		values := []string{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				values = append(values, string(char))
			}
		}

		if len(values) == 0 {
			panic("Len(values) is 0")
		} else if len(values) == 1 {
			val, _ := strconv.Atoi(fmt.Sprintf("%s%s", values[0], values[0]))
			sum += val
		} else {
			val, _ := strconv.Atoi(fmt.Sprintf("%s%s", values[0], values[len(values)-1]))
			sum += val
		}
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
