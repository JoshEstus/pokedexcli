package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		// extra white space but all lower
		{
			input:    "   hello   world  ",
			expected: []string{"hello", "world"},
		},
		// happy path
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		// mix of upper but no extra whitespace
		{
			input:    "HeLLo WoRlD",
			expected: []string{"hello", "world"},
		},
		// upper and extra whitespace
		{
			input:    "   HeLLo WoRlD   ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		/*
			Check the length of the actual slice against the expected slice
			if they don't match, use t.Errorf to print an error message
			and fail the test
		*/
		if len(actual) != len(c.expected) {
			t.Errorf("Actual Len and Expected Len do not match Actual %d != Expected %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual %s doesn't match Expected %s", word, expectedWord)
			}
		}

	}
}
