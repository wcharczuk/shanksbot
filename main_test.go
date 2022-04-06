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
		Input    string
		Expected bool
	}{
		{"", false},
		{"a", false},
		{"ab", false},
		{"aa", true},
		{"aaa", false},
		{"abaa", false},
		{"abab", true},
		{"abaab", false},
		{"abaaba", true},
		{"ababab", false},
		{"abababa", false},
		{"abcdabc", false},
		{"abcdabcd", true},
		{"abcdefghij", false},
		{"abcdefabcd", false},
		{"abcdeabcde", true},
	}

	for _, tc := range cases {
		if tc.Expected != repeats(tc.Input) {
			t.Errorf("fails repeats: %v expects %v", tc.Input, tc.Expected)
		}
	}
}
