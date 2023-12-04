package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"unicode"

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

func isSymbol(v rune) (bool, string) {
	if string(v) == "." || string(v) == " " {
		return false, ""
	}

	if unicode.IsNumber(v) {
		return false, ""
	}

	return true, string(v)
}

type NumberTracker struct {
	keep   bool
	symbol string
	str    strings.Builder
}

func part1(input string) int {
	board := parseInput(input)
	max_row := len(board) - 1
	max_col := len(board[0]) - 1
	fmt.Printf("max_row: %d, max_col:%d\n", max_row, max_col)

	checkAround := func(rr int, cc int) (bool, string) {
		hasSymbol := false
		symbol := "WTF"
		for _, roff := range []int{-1, 0, 1} {
			x := rr + roff
			if x < 0 || x >= max_row {
				continue
			}

			for _, coff := range []int{-1, 0, 1} {
				y := cc + coff
				if y < 0 || y > max_col {
					continue
				}
				retval, str := isSymbol(rune(board[x][y]))
				hasSymbol = hasSymbol || retval
				if retval {
					symbol = str
				}
			}
		}
		// fmt.Printf("s::%t\n", hasSymbol)
		return hasSymbol, symbol
	}

	// Rows
	sum := 0
	for ridx, row := range board {
		fmt.Printf("\n")
		t := NumberTracker{}
		for cidx, col := range row {
			// Skip all symbols
			b, _ := isSymbol(col)
			if t.str.Len() == 0 && b {
				// fmt.Printf("Skipping symbol\n")
				continue
			}

			if b && t.str.Len() > 0 {
				if t.keep {
					val := cast.ToInt(t.str.String())
					sum += val
					fmt.Printf("Adding %d(%s) to sum = %d\n", val, t.symbol, sum)
				}
				t.str.Reset()
				t.keep = false
				continue
			}

			// . without ongoing number
			if string(col) == "." && t.str.Len() == 0 {
				// fmt.Printf("Skipping dot - len(%d)\n", t.str.Len())
				continue
			}

			// . with ongoing number
			if string(col) == "." && t.str.Len() > 0 {
				if t.keep {
					val := cast.ToInt(t.str.String())
					sum += val
					fmt.Printf("Adding %d(%s) to sum = %d\n", val, t.symbol, sum)
				}
				t.str.Reset()
				t.keep = false
				// fmt.Printf("Reset - len(%d) %t\n", t.str.Len(), t.keep)
				continue
			}

			// Sanity check
			if !unicode.IsNumber(col) {
				panic("not symbol or number or .")
			}

			// Add current digit
			t.str.WriteRune(col)
			// fmt.Printf("Found digit: %s\n", t.str.String())
			b, symbol := checkAround(ridx, cidx)
			if b {
				t.symbol = symbol
			}
			t.keep = t.keep || b
			// Check for
			if t.keep && cidx == max_col {
				val := cast.ToInt(t.str.String())
				sum += val
				fmt.Printf("Adding %d to sum = %d\n", val, sum)

				t.str.Reset()
				t.keep = false
				// fmt.Printf("Found symbol on digit (%t)\n", t.keep)
			}

		}
	}

	return sum
}

func part2(input string) int {
	board := parseInput(input)
	max_row := len(board) - 1
	max_col := len(board[0]) - 1
	fmt.Printf("max_row: %d, max_col:%d\n", max_row, max_col)

	findNumber := func(r int, c int) string {
		number := string(board[r][c])
		// Remove the number not to add it in another iteration
		board[r] = util.ReplaceAtIndex(board[r], 'X', c)

		// check left and prepend digits
		for lx := c - 1; lx >= 0; lx-- {
			left := rune(board[r][lx])
			if !unicode.IsNumber(left) {
				break
			}
			number = string(left) + number
			// Remove the number not to add it in another iteration
			board[r] = util.ReplaceAtIndex(board[r], 'X', lx)
		}

		// Check right and append digits
		for rx := c + 1; rx <= max_col; rx++ {
			right := rune(board[r][rx])
			if !unicode.IsNumber(right) {
				break
			}
			number = number + string(right)
			// Remove the number not to add it in another iteration
			board[r] = util.ReplaceAtIndex(board[r], 'X', rx)
		}
		fmt.Printf("%s\n", number)
		return number
	}

	checkAround := func(rr int, cc int) []string {
		numbers := []string{}
		for _, roff := range []int{-1, 0, 1} {
			x := rr + roff
			if x < 0 || x > max_row {
				continue
			}

			for _, coff := range []int{-1, 0, 1} {
				y := cc + coff
				if y < 0 || y > max_col {
					continue
				}

				if unicode.IsNumber(rune(board[x][y])) {
					numbers = append(numbers, findNumber(x, y))
				}
			}
		}
		return numbers
	}

	// Rows
	sum := 0
	for ridx, row := range board {
		fmt.Printf("\n")
		for cidx, col := range row {
			if string(col) != "*" {
				continue
			}
			// Check around for number
			fmt.Printf("Found*\n")
			numbers := checkAround(ridx, cidx)
			if len(numbers) == 2 {
				sum += (cast.ToInt(numbers[0]) * cast.ToInt(numbers[1]))
				fmt.Printf("%s adding to sum: %d", numbers, sum)
			}
		}
	}

	for _, row := range board {
		fmt.Println(row)
	}
	return sum
}

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}
