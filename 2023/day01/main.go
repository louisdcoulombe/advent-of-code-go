package main

import (
	_ "embed"
	"errors"
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

func wordValue(values []string, line string) int {
	var (
		val int
		err error
	)
	if len(values) == 0 {
		panic("Len(values) is 0")
	} else if len(values) == 1 {
		ss := fmt.Sprintf("%s%s", values[0], values[0])
		val, err = strconv.Atoi(ss)
		if err != nil {
			panic("Error in Atoi #1")
		}
		fmt.Printf("%s : %s = %v\n", line, ss, val)
	} else {
		ss := fmt.Sprintf("%s%s", values[0], values[len(values)-1])
		val, err = strconv.Atoi(ss)
		if err != nil {
			panic("Error in Atoi #1")
		}
		fmt.Printf("%s : %s = %v\n", line, ss, val)
	}
	return val
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
		sum += wordValue(values, line)
	}

	return sum
}

func checkForNumberString(str string) (string, int, error) {
	// fmt.Printf("%s::", str)
	if len(str) < 3 {
		return "", 0, errors.New("too small")
	}

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, n := range numbers {
		if len(str) < len(n) {
			continue
		}
		if str[0:len(n)] == n {
			num := strconv.FormatInt(int64(i)+1, 10)
			// fmt.Printf("%s | ", num)
			return num, len(n) - 1, nil
		}
	}

	// fmt.Printf("\n---\n")
	return "", 0, errors.New("No number")
}

func part2(input string) int {
	parsed := parseInput(input)
	sum := 0
	// offset := 0
	for _, line := range parsed {
		values := []string{}
		for i, char := range line {
			// skip used chars
			//if offset > 0 {
			//offset -= 1
			//continue
			//}

			if unicode.IsDigit(char) {
				values = append(values, string(char))
				continue
			}

			// Check for string numbers
			var err error
			var v string
			v, _, err = checkForNumberString(line[i:])
			if err != nil {
				continue
			}
			values = append(values, v)
		}

		sum += wordValue(values, line)
	}
	return sum
}

func parseInput(input string) (ans []string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, line)
	}
	return ans
}
