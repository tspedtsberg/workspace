package main

import (
	"testing"
)


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
		input: "Hello world",
		expected: []string{"hello", "world"},
		}, 
		{
			input: "Today we're having funsies",
			expected: []string{"today", "we're", "having", "funsies"},
		},
		{
			input: "ButterFLY IS THA BEST POKEMON",
			expected: []string{"butterfly", "is", "tha", "best", "pokemon"},
		},
		{
			input: "test with 25 numbers",
			expected: []string{"test", "with", "25", "numbers"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match: %v vs %v", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}

	}

}