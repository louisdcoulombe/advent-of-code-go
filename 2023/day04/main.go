package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/louisdcoulombe/advent-of-code-go/cast"
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

func sortedIntList(cards []string) []int {
	list := []int{}
	for _, c := range cards {
		// Deal with 1 digit number in fixed double spacing
		if len(c) == 0 {
			continue
		}
		list = append(list, cast.ToInt(c))
	}
	sort.Ints(list)
	return list
}

func gameCount(gameCards *[]int, handCards *[]int) int {
	count := 0
	for _, gc := range *gameCards {
		for _, hc := range *handCards {
			// Sorted list, if already greater it's not there
			if hc > gc {
				break
			}

			// Found same card
			if hc == gc {
				count += 1
				break
			}
		}
	}
	return count
}

func part1(input string) int {
	parsed := parseInput(input)
	sum := 0
	for _, line := range parsed {
		// game card | your hand
		parts := strings.Split(line, "|")
		parts[0] = strings.TrimSpace(strings.Split(parts[0], ":")[1])
		gameCards := sortedIntList(strings.Split(parts[0], " "))
		handCards := sortedIntList(strings.Split(parts[1], " "))
		count := gameCount(&gameCards, &handCards)
		if count > 0 {
			sum += 1 << (count - 1)
		}
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	cards := make([]int, len(parsed))
	for idx, line := range parsed {
		// game card | your hand
		parts := strings.Split(line, "|")
		parts[0] = strings.TrimSpace(strings.Split(parts[0], ":")[1])
		gameCards := sortedIntList(strings.Split(parts[0], " "))
		handCards := sortedIntList(strings.Split(parts[1], " "))

		count := gameCount(&gameCards, &handCards)
		replay := 0
		for replay < cards[idx]+1 {
			// Increment next cards counts
			for i := 0; i < count; i++ {
				// dont overflow
				if (i + idx + 1) > len(parsed)-1 {
					break
				}
				cards[i+idx+1] += 1
			}
			replay++
		}
		// Record the original copy
		cards[idx] += 1

		// fmt.Printf("%d(%d):: %d\n", idx+1, count, cards)
	}
	return util.Sum[int, int](cards)
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
