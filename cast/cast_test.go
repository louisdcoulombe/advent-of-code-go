package cast

import (
	"testing"
)

func Test_cast(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example_1",
			input: "12",
			want:  12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.input); got != tt.want {
				t.Errorf("cast.ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cast_rune(t *testing.T) {
	tests := []struct {
		name  string
		input byte
		want  int
	}{
		{
			name:  "example_1",
			input: "12"[0],
			want:  1,
		},
		{
			name:  "example_1",
			input: "12"[1],
			want:  2,
		},
		{
			name:  "example_1",
			input: "0"[0],
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.input); got != tt.want {
				t.Errorf("cast.ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
