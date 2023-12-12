package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	// "github.com/louisdcoulombe/advent-of-code-go/cast"
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

const (
	LEFT  = 0
	RIGHT = 1
)

type RingBuffer struct {
	next   int
	buffer []int
}

func (b *RingBuffer) Next() (next int) {
	next = b.buffer[b.next]
	b.next = (b.next + 1) % len(b.buffer)
	return next
}

func (b *RingBuffer) FromString(s string) {
	b.next = 0

	for _, c := range s {
		if string(c) == "L" {
			b.buffer = append(b.buffer, LEFT)
		} else {
			b.buffer = append(b.buffer, RIGHT)
		}
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
