package main

import (
	"reflect"
	"testing"
)

func TestNormalizeAndSplitInput(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "lowercase conversion",
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			name:     "extra whitespace removal",
			input:    "  multiple    spaces  \t and tabs  ",
			expected: []string{"multiple", "spaces", "and", "tabs"},
		},
		{
			name:     "empty input",
			input:    "   ",
			expected: []string{}, // Or nil, depending on your goal
		},
		{
			name: 	  "test white space",
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			name: 	  "test white space and capitals",
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	// Run each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := NormalizeAndSplitInput(tc.input)
			
			// reflect.DeepEqual is used to compare slices in Go
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %q, but got %q", tc.expected, actual)
			}
		})
	}
}