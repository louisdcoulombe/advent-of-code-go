package main

import (
	"testing"
)

var (
	example1 = ``
	example2 = ``
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		//
		{
			name:  "rn=1",
			input: "rn=1",
			want:  30,
		},
		//
		{
			name:  "cm-",
			input: "cm-",
			want:  253,
		},
		//
		{
			name:  "qp=3",
			input: "qp=3",
			want:  97,
		},
		//
		{
			name:  "cm=2",
			input: "cm=2",
			want:  47,
		},
		//
		{
			name:  "qp-",
			input: "qp-",
			want:  14,
		},
		//
		{
			name:  "pc=4",
			input: "pc=4",
			want:  180,
		},
		//
		{
			name:  "ot=9",
			input: "ot=9",
			want:  9,
		},
		//
		{
			name:  "ab=5",
			input: "ab=5",
			want:  197,
		},
		//
		{
			name:  "pc-",
			input: "pc-",
			want:  48,
		},
		//
		{
			name:  "pc=6",
			input: "pc=6",
			want:  214,
		},
		//
		{
			name:  "ot=7",
			input: "ot=7",
			want:  231,
		},
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
