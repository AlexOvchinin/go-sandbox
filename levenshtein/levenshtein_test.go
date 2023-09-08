package levenshtein

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"Anne", "Ann", 1},
		{"Ane", "Ann", 1},
		{"An", "Ann", 1},
		{"A", "Ann", 2},
		{"Aen", "Ann", 1},
		{"Aen", "Anne", 2},
		{"Ann", "Ann", 0},
	}

	for _, c := range cases {
		got := Distance(c.a, c.b)
		if got != c.want {
			t.Errorf("Distance(%v, %v) == %v, but expected %v", c.a, c.b, got, c.want)
		}
	}
}
