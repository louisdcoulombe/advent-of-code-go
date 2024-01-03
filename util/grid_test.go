package util

import (
	// "fmt"
	"sort"
	"strings"
	"testing"
)

func parseInput(input string) (ans []string) {
	return append(ans, strings.Split(input, "\n")...)
}

func Test_Append(t *testing.T) {
	g := GridRow{}
	g = append(g, NewPoint(0, 0))
	if len(g) == 0 {
		t.Errorf("Append doesnt work")
	}
}

func Test_Contains(t *testing.T) {
	p := NewPoint(0, 1)
	want := GridRow{NewPoint(0, 1), NewPoint(0, 0)}
	if !want.Contains(p) {
		t.Errorf("Not in!")
	}

	p = NewPoint(0, 0)
	if !want.Contains(p) {
		t.Errorf("0,0 Not in!")
	}
}

func Test_IsIn(t *testing.T) {
	p := NewPoint(0, 1)
	want := GridRow{NewPoint(0, 1), NewPoint(0, 0)}
	if !p.IsIn(want) {
		t.Errorf("Not in!")
	}

	p = NewPoint(0, 0)
	if !p.IsIn(want) {
		t.Errorf("0,0 Not in!")
	}
}

var example1 = `abcdei
fghij
klmno
pqrst
uvwxy
`

func Test_CheckNeighbours(t *testing.T) {
	tests := []struct {
		name    string
		input   GridPoint
		wants   GridRow
		corners bool
	}{
		{
			name:    "Top left",
			input:   NewPoint(0, 0),
			wants:   GridRow{NewPointV(1, 0, "b"), NewPointV(0, 1, "f"), NewPointV(1, 1, "g")},
			corners: true,
		},
		{
			name:    "Top right",
			input:   NewPoint(4, 0),
			wants:   GridRow{NewPointV(3, 0, "d"), NewPointV(3, 1, "i"), NewPointV(4, 1, "j")},
			corners: true,
		},
		{
			name:    "Bottom left",
			input:   NewPoint(0, 4),
			wants:   GridRow{NewPointV(0, 3, "p"), NewPointV(1, 3, "q"), NewPointV(1, 4, "v")},
			corners: true,
		},
		{
			name:    "Bottom right",
			input:   NewPoint(4, 4),
			wants:   GridRow{NewPointV(4, 3, "t"), NewPointV(3, 3, "s"), NewPointV(3, 4, "x")},
			corners: true,
		},
		{
			name:  "Center m",
			input: NewPoint(2, 2),
			wants: GridRow{
				NewPointV(1, 1, "g"), NewPointV(2, 1, "h"), NewPointV(3, 1, "i"),
				NewPointV(1, 2, "l"), NewPointV(3, 2, "n"),
				NewPointV(1, 3, "q"), NewPointV(2, 3, "r"), NewPointV(3, 3, "s"),
			},
			corners: true,
		},

		{
			name:    "Top left no-c",
			input:   NewPoint(0, 0),
			wants:   GridRow{NewPointV(1, 0, "b"), NewPointV(0, 1, "f")},
			corners: false,
		},
		{
			name:    "Top right no-c",
			input:   NewPoint(4, 0),
			wants:   GridRow{NewPointV(3, 0, "d"), NewPointV(4, 1, "j")},
			corners: false,
		},
		{
			name:    "Bottom left no-c",
			input:   NewPoint(0, 4),
			wants:   GridRow{NewPointV(0, 3, "p"), NewPointV(1, 4, "v")},
			corners: false,
		},
		{
			name:    "Bottom right no-c",
			input:   NewPoint(4, 4),
			wants:   GridRow{NewPointV(4, 3, "t"), NewPointV(3, 4, "x")},
			corners: false,
		},
		{
			name:  "Center m no-c",
			input: NewPoint(2, 2),
			wants: GridRow{
				NewPointV(2, 1, "h"),
				NewPointV(1, 2, "l"), NewPointV(3, 2, "n"),
				NewPointV(2, 3, "r"),
			},
			corners: false,
		},
	}

	for _, tt := range tests {
		g := MakeGrid(parseInput(example1))
		gots := g.GetNeighbours(tt.input, tt.corners)
		// fmt.Printf("%v", gots)
		// fmt.Printf("Test '%s'::", tt.name)
		// for _, s := range g.grid {
		// 	fmt.Printf("%s\n", s)
		// }

		sort.Sort(GridRow(gots))
		sort.Sort(GridRow(tt.wants))

		if len(gots) != len(tt.wants) {
			t.Errorf("'%s' -> GetNeighbours(%v) = %v, want %v", tt.name, tt.input, gots, tt.wants)
			continue
		}

		for i, want := range tt.wants {
			if !gots[i].Equal(want) {
				t.Errorf("'%s' -> GetNeighbours(%v) = %v, want %v (%d)", tt.name, tt.input, gots, tt.wants, i)
			}
		}
	}
}
