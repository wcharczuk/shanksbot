package main

import "testing"

func Test_primeReciprocalRepeatsAfter(t *testing.T) {
	cases := [...]struct {
		Input    int
		Expected int
	}{
		{7, 6},
		{60013, 5001},
		{60017, 60016},
		{61561, 405},
		{61583, 61582},
		{62039, 31019},
	}

	for _, tc := range cases {
		actual := primeReciprocalRepeatsAfter(tc.Input)
		if tc.Expected != actual {
			t.Errorf("input: %d expects: %d, actual: %d", tc.Input, tc.Expected, actual)
		}
	}
}

func Test_repeats(t *testing.T) {
	cases := [...]struct {
		Input    []int
		Expected bool
	}{
		{nil, false},
		{[]int{}, false},
		{[]int{0}, false},
		{[]int{0, 1}, false},
		{[]int{0, 0}, true},
		{[]int{0, 0, 0}, false},
		{[]int{0, 1, 0, 0}, false},
		{[]int{0, 1, 0, 1}, true},
		{[]int{0, 1, 0, 0, 1}, false},
		{[]int{0, 1, 0, 0, 1, 0}, true},
		{[]int{0, 1, 0, 1, 0, 1}, false},
		{[]int{0, 1, 0, 1, 0, 1, 0}, false},
		{[]int{0, 1, 0, 1, 0, 1, 2}, false},
		{[]int{0, 1, 2, 3, 0, 1, 2, 3}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, false},
	}

	for _, tc := range cases {
		if tc.Expected != repeats(tc.Input) {
			t.Errorf("fails repeats: %v expects %v", tc.Input, tc.Expected)
		}
	}
}
