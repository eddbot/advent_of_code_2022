package main

import "testing"

type TableTest struct {
	input string
	want  int
}

func TestPart1(t *testing.T) {

	tests := []TableTest{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 7},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
	}

	for _, tt := range tests {
		got := solver(tt.input, 4)
		want := tt.want

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {

	tests := []TableTest{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 19},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 23},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", want: 23},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 29},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 26},
	}

	for _, tt := range tests {
		got := solver(tt.input, 14)
		want := tt.want

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
