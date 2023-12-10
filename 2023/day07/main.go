package main

import (
	// "container/heap"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/louisdcoulombe/advent-of-code-go/cast"
	"github.com/louisdcoulombe/advent-of-code-go/util"
)

func PrintMap(m map[string]int) {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

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

func FindSimilar(hand string, card string) int {
	count := 0
	for _, r := range hand {
		if string(r) == card {
			count++
		}
	}
	return count
}

func ParseHand(hand string) map[int]int {
	m := map[int]int{}
	for i, c := range hand {
		key := CARDS[string(c)]
		if val, ok := m[key]; ok {
			m[key] = val + FindSimilar(string(hand[(i+1):]), string(c))
		} else {
			m[key] = FindSimilar(string(hand[(i+1):]), string(c))
		}
	}
	// fmt.Printf("%v\n", m)
	return m
}

var CARDS = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type Hand struct {
	hand   string
	bet    int
	counts map[int]int
	power  int // 0 (highcard) to 6(5 of a kind)
}

func (h *Hand) CalculatePower() {
	// 1(0) : high card
	// 2(1) : pair
	// 4(2+2) : 2 pair
	// 8(3) : 3ofakind
	// 16(3+1) : fullhouse
	// 32(4) : quad
	// 64(5) : 5ofakind
	for _, v := range h.counts {
		// fmt.Printf("%d=%v\n", k, v)
		h.power += v
	}
}

func (h *Hand) UpdateJacks() {
	max_card := func(m map[int]int) int {
		max_card := 0
		max_val := 0
		for k, v := range m {
			if v > max_val {
				max_card = k
				max_val = v
			}
		}
		return max_card
	}

	countJacks := func(m map[int]int) int {
		for k, v := range m {
			if k == 11 {
				return v
			}
		}
		return 0
	}

	nbJack := countJacks(h.counts)
	if nbJack == 0 {
		return
	}

	maxCard := max_card(h.counts)
	h.counts[maxCard] += nbJack
	h.counts[11] = 0
}

type HandHeap []Hand

func (h HandHeap) Len() int      { return len(h) }
func (h HandHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h HandHeap) Less(i, j int) bool {
	// Type is greater
	if h[i].power < h[j].power {
		fmt.Printf("%v < %v\n", h[i], h[j])
		return true
	}
	// Type is less
	if h[i].power > h[j].power {
		fmt.Printf("%v > %v\n", h[i], h[j])
		return false
	}
	//
	// max_card := func(m map[int]int) int {
	// 	max_card := 0
	// 	max_val := 0
	// 	for k, v := range m {
	// 		if v > max_val {
	// 			max_card = k
	// 			max_val = v
	// 		}
	// 	}
	// 	return max_card
	// }
	//
	// max_i := max_card(h[i].counts)
	// max_j := max_card(h[j].counts)
	// if max_i < max_j {
	// 	return true
	// }
	// if max_i > max_j {
	// 	return false
	// }
	//
	for idx := range h[i].hand {
		left := CARDS[string(h[i].hand[idx])]
		right := CARDS[string(h[j].hand[idx])]
		if left < right {
			fmt.Printf("c: %v(%d) < %v)%d\n", h[i], left, h[j], right)
			return true
		}
		if left > right {
			fmt.Printf("c: %v(%d) > %v)%d\n", h[i], left, h[j], right)
			return false
		}
	}

	// Same type, check cards
	fmt.Printf("%v == %v", h[i], h[j])
	panic("Exact same value!")
}

func (h *HandHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Hand))
}

func (h *HandHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func part1(input string) int {
	parsed := parseInput(input)
	hands := []Hand{}
	for _, line := range parsed {
		parts := strings.Split(line, " ")
		co := ParseHand(parts[0])
		h := Hand{
			hand:   parts[0],
			bet:    cast.ToInt(parts[1]),
			counts: co,
		}
		h.CalculatePower()
		hands = append(hands, h)

		// fmt.Printf("%v\n", h)
	}

	sort.Sort(HandHeap(hands))

	sum := 0
	for rank, hand := range hands {
		score := hand.bet * (rank + 1)
		sum += hand.bet * (rank + 1)
		fmt.Printf("p:%d | %v > %d = %d\n", hand.power, hand, score, sum)
		// fmt.Printf("%v > %d = %d\n", hand, score, sum)
	}
	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	hands := []Hand{}
	for _, line := range parsed {
		parts := strings.Split(line, " ")
		co := ParseHand(parts[0])
		h := Hand{
			hand:   parts[0],
			bet:    cast.ToInt(parts[1]),
			counts: co,
		}
		h.UpdateJacks()
		h.CalculatePower()
		hands = append(hands, h)

		// fmt.Printf("%v\n", h)
	}

	sort.Sort(HandHeap(hands))

	sum := 0
	for rank, hand := range hands {
		score := hand.bet * (rank + 1)
		sum += hand.bet * (rank + 1)
		fmt.Printf("p:%d | %v > %d = %d\n", hand.power, hand, score, sum)
		// fmt.Printf("%v > %d = %d\n", hand, score, sum)
	}
	return sum
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
