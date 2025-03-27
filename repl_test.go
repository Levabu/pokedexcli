package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "  heLLo      wOrld  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		la := len(actual)
		le := len(c.expected)
		if le != la {
			t.Errorf("Slices' sizes don't match: \n\texpected – %v, \n\tactual – %v", le, la)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("Words at index %v don't match: \n\texpected – %v, \n\tactual – %v", i, expectedWord, word)
			}
		}
	}
}