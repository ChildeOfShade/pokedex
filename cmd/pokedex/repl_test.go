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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "\tSquirtle\nMewtwo ",
			expected: []string{"squirtle", "mewtwo"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// First, check that lengths match
		if len(actual) != len(c.expected) {
			t.Errorf("for input %q, expected %d words, got %d", c.input, len(c.expected), len(actual))
			continue
		}

		// Compare word-by-word
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("for input %q, expected %v but got %v", c.input, c.expected, actual)
				break
			}
		}
	}
}
