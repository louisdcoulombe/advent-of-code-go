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

func Hash(in string) int {
	currentValue := 0
	for _, char := range in {
		currentValue += int(char)
		currentValue = (currentValue * 17) % 256
	}
	return currentValue
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
	inputs := strings.Split(parsed[0], ",")
	sum := 0
	for _, input := range inputs {
		sum += Hash(input)
	}

	return sum
}

func RemoveItems(removed []int, box *Box) {
	for _, i := range removed {
		length := len(*box)
		idx_min := max(0, i)
		idx_max := min(length-1, i+1)
		*box = append((*box)[0:idx_min], (*box)[i:idx_max]...)
	}
}

func cmd_dash(boxes *Boxes, in Lens) {
	// Calculate hash
	hash := Hash(in.op)
	box := &(*boxes)[hash]
	// if len exist in box = hash
	removed := []int{}
	for i, b := range *box {
		if b.name == in.name {
			removed = append(removed, i)
		}
	}

	// remove the lens from box
	// move other lens behind forward
}

func cmd_equal(boxes *Boxes, in Lens) int {
	// Calculate hash
	// _ := Hash(in.op)
	// Extract focus
	// Remove len of same type if exist
	// Otherwise add at the end
	return 0
}

type Lens struct {
	op    string
	name  string
	focus int
}
type (
	Box   []Lens
	Boxes [][]Lens
)

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
