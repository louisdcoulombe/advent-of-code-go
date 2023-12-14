package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

func (b *RingBuffer) Reset() {
	b.next = 0
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

type Movements map[string][]string

func ParseMovements(in []string) Movements {
	re := regexp.MustCompile(`([1-9A-Z]{3}) = \(([1-9A-Z]{3}), ([1-9A-Z]{3})\)`)
	m := Movements{}
	for i, l := range in {
		if i <= 1 || l == "" {
			continue
		}

		matches := re.FindStringSubmatch(l)
		// fmt.Printf("%v > ", matches)
		if len(matches) < 4 {
			panic(fmt.Sprintf("%d:%s=[%v]\n", i, l, matches))
		}
		// fmt.Printf("[%s]: (%s, %s)\n", matches[1], matches[2], matches[3])
		m[matches[1]] = []string{matches[2], matches[3]}
	}
	return m
}

func EndsWith(s string, end string) bool {
	return string(s[len(s)-1]) == end
}

type Movement struct {
	key  string
	done bool
}

func MovementsKeysEndsWith(m Movements, end string) (ans []Movement) {
	for k := range m {
		if EndsWith(k, end) {
			ans = append(ans, Movement{k, false})
		}
	}
	return ans
}

func part1(input string) int {
	parsed := ParseInput(input)
	rb := RingBuffer{}
	rb.FromString(parsed[0])
	m := ParseMovements(parsed)
	// fmt.Printf("%v\n", m)

	count := 0
	dst := "AAA"
	for dst != "ZZZ" {
		next := rb.Next()
		arr, err := m[dst]
		// fmt.Printf("%d| %d-%s->%v\n", count, next, dst, arr)
		if !err {
			panic("Ouch")
		}
		dst = arr[next]
		count++
	}

	return count
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func FindCycles(rb RingBuffer, m Movements) (ans []int) {
	initals := MovementsKeysEndsWith(m, "A")
	for _, dst := range initals {
		rb.Reset()
		count := 0
		for !dst.done {
			next := rb.Next()
			arr, ok := m[dst.key]
			if !ok {
				panic("Ouch")
			}
			dst.key = arr[next]
			dst.done = EndsWith(dst.key, "Z")
			count++
		}

		ans = append(ans, count)
	}

	return ans
}

func part2(input string) int {
	parsed := ParseInput(input)
	rb := RingBuffer{}
	rb.FromString(parsed[0])
	fmt.Printf("%v\n\n", rb)

	m := ParseMovements(parsed)
	cycles := FindCycles(rb, m)
	return LCM(cycles[0], cycles[1], cycles[2:]...)
}

func ParseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
