package main

import (
	"fmt"
	"testing"
)

var (
	example0 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`
	example1 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
	example2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`
)

func Test_RB(t *testing.T) {
	r := RingBuffer{}
	r.FromString("RL")
	ans := []int{}
	for i := 0; i < 4; i++ {
		ans = append(ans, r.Next())
	}

	expected := []int{RIGHT, LEFT, RIGHT, LEFT}
	for i, c := range ans {
		if c != expected[i] {
			t.Errorf("%v != %v, index:%d", ans, expected, i)
		}
	}
}

func Test_Movements(t *testing.T) {
	m := ParseMovements(ParseInput(example0))
	fmt.Printf("%v", m)
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example0",
			input: example0,
			want:  2,
		},
		{
			name:  "example1",
			input: example1,
			want:  6,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example2",
			input: example2,
			want:  6,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
