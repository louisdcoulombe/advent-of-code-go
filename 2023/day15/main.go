package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

func cmd_dash(boxes *Boxes, in Lens) {
	box := make(Box, len((*boxes)[in.hash]))
	copy(box, (*boxes)[in.hash])
	// fmt.Printf("%d: %v\n", in.hash, (*boxes)[in.hash])

	(*boxes)[in.hash] = Box{}
	// if len exist in box = hash
	for _, b := range box {
		// if the lens exist, remove it
		if b.name == in.name {
			continue
		}
		// fmt.Printf("[- %v]", b)
		(*boxes)[in.hash] = append((*boxes)[in.hash], b)
	}

	// fmt.Printf("- > %v\n", (*boxes)[in.hash])
}

func cmd_equal(boxes *Boxes, in Lens) {
	// Calculate hash
	hash := in.hash
	box := &(*boxes)[hash]
	// Extract focus
	if len(*box) == 0 {
		*box = append(*box, in)
		// fmt.Printf("- > %v\n", (*boxes)[hash])
		return
	}

	for i, b := range *box {
		if b.name == in.name {
			(*box)[i] = in
			// fmt.Printf("- > %v\n", (*boxes)[hash])
			return
		}
	}

	*box = append(*box, in)
	// fmt.Printf("- > %v\n", (*boxes)[hash])
}

type Lens struct {
	isEqual bool
	hash    int
	op      string
	name    string
	focus   int
}
type (
	Box   []Lens
	Boxes [][]Lens
)

func FromOp(op string) Lens {
	l := Lens{}
	l.op = op
	if strings.Contains(op, "=") {
		l.isEqual = true
		str := strings.Split(op, "=")
		l.name = str[0]
		var ok error
		l.focus, ok = strconv.Atoi(str[1])
		if ok != nil {
			panic("Atoi failed")
		}
	} else {
		l.isEqual = false
		l.name = strings.Split(op, "-")[0]
		l.focus = 0
	}
	l.hash = Hash(l.name)
	// fmt.Printf("%v", l)
	return l
}

func part2(input string) int {
	boxes := make(Boxes, 256)

	parsed := parseInput(input)
	for _, op := range strings.Split(parsed[0], ",") {
		lens := FromOp(op)
		if lens.isEqual {
			cmd_equal(&boxes, lens)
		} else {
			cmd_dash(&boxes, lens)
		}
	}

	sum := 0
	for i, box := range boxes {
		for j, lens := range box {
			sum += (1 + i) * ((1 + j) * lens.focus)
			// fmt.Printf("(%d, %d) %v = %d\n", i, j, lens, sum)
		}
	}
	return sum
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
