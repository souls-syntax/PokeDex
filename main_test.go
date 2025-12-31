package main

import (
	"testing"
	"reflect"
)

func TestCleanInput(t *testing.T) {
	type test struct {
		input string
		expected []string
	}

	tests := []test{
		{
			input: "  Hello  world  ",
			expected: []string{"hello","world"},
		},
		{
			input: "hello      world",
			expected: []string{"hello", "world"},
		},
	}
	for i, c := range tests {
		got := cleanInput(c.input)
		if !reflect.DeepEqual(c.expected, got) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1,c.expected,got)
		}
	}
}
