package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"sync"

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
	allZero := func() bool {
		for _, i := range h {
			if i != 0 {
				return false
			}
		}
		return true
	}

	if allZero() {
		return 0
	}

	return h[len(h)-1] + extrapolate(calculateDiffs(h))
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	histories := parseHistories(parsed)

	sum := 0
	for _, h := range histories {
		sum += extrapolate(h)
	}
	return sum
}

func parallel_part1(input string) int {
	parsed := parseInput(input)
	_ = parsed
	histories := parseHistories(parsed)
	// fmt.Printf("%v", histories)

	numJobs := len(histories)
	jobs := make(chan History, len(histories))
	results := make(chan int, len(histories))

	worker := func(input <-chan History, result chan<- int) {
		for h := range input {
			result <- extrapolate(h)
		}
	}
	for w := 0; w < 8; w++ {
		go worker(jobs, results)
	}

	for _, h := range histories {
		jobs <- h
	}
	close(jobs)

	// Fan-in: Collect results
	var wg sync.WaitGroup
	wg.Add(numJobs) // Set WaitGroup counter to the number of jobs

	// Launch a goroutine to wait for all jobs to finish
	go func() {
		wg.Wait()      // Wait for all jobs to be done
		close(results) // Close the results channel after all jobs are processed
	}()

	// Process results
	sum := 0
	for result := range results {
		sum += result
		wg.Done() // Decrease the WaitGroup counter as each result is processed
	}

	return sum
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
