package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/louisdcoulombe/advent-of-code-go/cast"
	datastructures "github.com/louisdcoulombe/advent-of-code-go/data-structures"
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

func parseMaps(raw []string) (mappings []datastructures.MappingList) {
	singleMap := datastructures.MappingList{}
	sortMap := func() {
		sort.Slice(singleMap, func(i, j int) bool {
			return singleMap[i].Src < singleMap[j].Src
		})
		// for _, m := range singleMap {
		// fmt.Printf("%d:%d->%d:%d [%d]\n", m.Src, m.Src_max, m.Dst, m.Dst_max, m.Count)
		// }
	}

	// currentMapping := ""
	for idx, line := range raw {
		// skip first section
		if idx < 2 {
			continue
		}
		// Ending of mapping
		if strings.TrimSpace(line) == "" {
			sortMap()
			mappings = append(mappings, singleMap)
			singleMap = []datastructures.Mapping{}
			continue
		}
		// New mapping
		if strings.Contains(line, "map") {
			// currentMapping = strings.Split(line, " ")[0]
			// fmt.Println(currentMapping)
			continue
		}

		parts := strings.Split(line, " ")
		m := datastructures.Mapping{}
		m.Dst = cast.ToInt(parts[0])
		m.Src = cast.ToInt(parts[1])
		m.Count = cast.ToInt(parts[2])
		m.Src_max = m.Src + m.Count
		m.Dst_max = m.Dst + m.Count
		singleMap = append(singleMap, m)

		// End of file
		if idx == len(raw)-1 {
			sortMap()
			mappings = append(mappings, singleMap)
		}
	}
	return mappings
}

func findLocation(maps []datastructures.MappingList, seed int) int {
	// fmt.Printf("Seed: %d\n", seed)

	var next int = seed
	for _, mapList := range maps {
		// fmt.Printf("%d->", next)
		for _, m := range mapList {
			// Find dst otherwise leave it there
			if m.Contains(next) {
				next = m.Get(next)
				break
			}
		}
		// fmt.Printf("%d\n", next)
	}
	return next
}

func part1(input string) int {
	parsed := parseInput(input)
	maps := parseMaps(parsed)

	seeds := strings.Split(strings.Split(parsed[0], ":")[1], " ")[1:]
	// fmt.Println(seeds)
	min_location := math.MaxInt64
	for _, s := range seeds {
		seed := cast.ToInt(strings.TrimSpace(s))
		next := findLocation(maps, seed)
		if next < min_location {
			min_location = next
		}
	}

	return min_location
}

func part2(input string) int {
	parsed := parseInput(input)
	maps := parseMaps(parsed)

	expandSeed := func(s int, r int) (seeds []int) {
		for i := 0; i < r; i++ {
			seeds = append(seeds, s+i)
		}
		return seeds
	}

	seeds := strings.Split(strings.Split(parsed[0], ":")[1], " ")[1:]
	numJobs := len(seeds) / 2
	results := make(chan int, numJobs)

	worker := func(seed int, offset int, results chan<- int) {
		min_location := math.MaxInt64
		for _, s := range expandSeed(seed, offset) {
			next := findLocation(maps, s)
			if next < min_location {
				min_location = next
			}
		}

		results <- min_location
	}

	// fmt.Println(seeds)
	for i := 0; i < len(seeds); i = (i + 2) {
		seed := cast.ToInt(strings.TrimSpace(seeds[i]))
		offset := cast.ToInt(strings.TrimSpace(seeds[i+1]))
		go worker(seed, offset, results)
	}

	// Fan-in: Collect results
	var wg sync.WaitGroup
	wg.Add(numJobs) // Set WaitGroup counter to the number of jobs

	// Launch a goroutine to wait for all jobs to finish
	go func() {
		wg.Wait()      // Wait for all jobs to be done
		close(results) // Close the results channel after all jobs are processed
	}()

	// Process results
	location := math.MaxInt64
	for result := range results {
		if result < location {
			location = result
		}
		wg.Done() // Decrease the WaitGroup counter as each result is processed
	}

	return location
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
