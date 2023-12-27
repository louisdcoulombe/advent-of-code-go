package main

import (
	"testing"
)

var (
	example1 = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
	example2 = example1
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  1320,
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

func Test_Hash(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "HASH",
			input: "HASH",
			want:  52,
		},
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
			if got := Hash(tt.input); got != tt.want {
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
			want:  145,
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
