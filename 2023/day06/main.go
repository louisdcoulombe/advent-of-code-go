package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func StringsToInts(str string, sep string) (result []int) {
	parts := strings.Split(str, sep)
	for _, s := range parts {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		result = append(result, cast.ToInt(s))
	}
	return result
}

type Round struct {
	time     int
	distance int
}

func FindWinners(rounds []Round) (winners int) {
	winners = 0
	for r, round := range rounds {
		winner := 0
		for time := 0; time < round.time; time++ {
			dist := time * (round.time - time)
			if dist > round.distance {
				winner++
				// fmt.Printf("W(%d) push:%d - distance:%d\n", winners, time, dist)
			}
		}

		if r == 0 {
			winners = winner
		} else {
			winners = winners * winner
		}

	}

	return winners
}

func part1(input string) int {
	parsed := parseInput(input)
	timeStr := strings.Split(parsed[0], ":")[1]
	distanceStr := strings.Split(parsed[1], ":")[1]
	times := StringsToInts(timeStr, " ")
	distances := StringsToInts(distanceStr, " ")
	rounds := []Round{}
	for i := range times {
		rounds = append(rounds, Round{times[i], distances[i]})
	}
	return FindWinners(rounds)
}

func part2(input string) int {
	parsed := parseInput(input)
	time := cast.ToInt(strings.ReplaceAll(strings.Split(parsed[0], ":")[1], " ", ""))
	distance := cast.ToInt(strings.ReplaceAll(strings.Split(parsed[1], ":")[1], " ", ""))
	// fmt.Printf("t: %d - d:%d\n", time, distance)

	return FindWinners([]Round{{time, distance}})
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		// fmt.Printf("%s\n", line)
		ans = append(ans, line)
	}
	return ans
}
