package main

import (
	"testing"
)

var example_1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var example_2 = ``

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example_1,
			want:  142,
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
			name:  "example",
			input: example_2,
			want:  0,
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
