package main

import "testing"

func TestRobCircle(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 3, 2}, 3}, // first and last can't both be taken
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
		{[]int{5}, 5},
		{[]int{}, 0},
		{[]int{2, 1}, 2},
		{[]int{200, 3, 140, 20, 10}, 340},
	}
	for _, c := range cases {
		if got := RobCircle(c.in); got != c.want {
			t.Errorf("RobCircle(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
