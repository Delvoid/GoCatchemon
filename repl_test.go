package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"help", []string{"help"}},
		{"exit", []string{"exit"}},
		{"help exit", []string{"help", "exit"}},
		{"help exit help", []string{"help", "exit", "help"}},
		{"help  ", []string{"help"}},
		{"  help", []string{"help"}},
		{"  help  ", []string{"help"}},
		{"  help  exit  ", []string{"help", "exit"}},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		if len(got) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %q, want %q", c.input, got, c.expected)
		}

		for i := range got {
			if got[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) == %q, want %q", c.input, got, c.expected)
			}
		}

	}

}
