package golang

import (
	"slices"
	"testing"
)

func TestUnique(t *testing.T) {
	for i, ss := range []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 5, 0, 10, 3, 5, 0, 4, 8, 2, 9, 1, 6, 7, 10, 8, 4, 1, 9, 0, 5, 5, 5, 5, 5},
			want:  []int{1, 5, 0, 10, 3, 4, 8, 2, 9, 6, 7},
		},
	} {
		if got := Unique(ss.input); !slices.Equal(got, ss.want) {
			t.Errorf("%d. Unique(%v): want %v, got %v", i+1, ss.input, ss.want, got)
		}
	}
}

func TestSlices(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, r := range []struct {
		length int
		want   []int
	}{
		{-1, []int{}},
		{0, []int{}},
		{5, []int{0, 1, 2, 3, 4}},
		{10, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{15, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4}},
	} {
		if got := Slices(r.length, s...); !slices.Equal(got, r.want) {
			t.Errorf("%d. Slices(%d, %v): want %v, got %v", i+1, r.length, s, r.want, got)
		}
	}
}
