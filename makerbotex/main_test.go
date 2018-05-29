package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxMinDifference(t *testing.T) {
	type TestCase struct {
		Input    []string
		Expected int
	}

	tests := map[string]TestCase{
		"Only 1 element": {
			Input:    []string{"5"},
			Expected: 0,
		},
		"Sorted list": {
			Input:    []string{"1", "2", "3", "4", "5"},
			Expected: 4,
		},
		"Unsorted list": {
			Input:    []string{"5", "3", "2", "8", "1"},
			Expected: 7,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Expected, getMaxMinDifference(test.Input))
		})
	}
}

func TestStrMinusCharAtIndex(t *testing.T) {
	assert.Equal(t, "", strMinusCharAtIndex("A", 0))
	assert.Equal(t, "B", strMinusCharAtIndex("AB", 0))
	assert.Equal(t, "BC", strMinusCharAtIndex("ABC", 0))
	assert.Equal(t, "AC", strMinusCharAtIndex("ABC", 1))
	assert.Equal(t, "AB", strMinusCharAtIndex("ABC", 2))
	assert.Equal(t, "ABDE", strMinusCharAtIndex("ABCDE", 2))
}

func TestPermutations(t *testing.T) {
	type TestCase struct {
		Input    string
		Expected []string
	}

	tests := map[string]TestCase{
		"only 1 character": {
			Input:    "1",
			Expected: []string{"1"},
		},
		"2 characters": {
			Input:    "12",
			Expected: []string{"12", "21"},
		},
		"3 characters": {
			Input:    "123",
			Expected: []string{"123", "132", "213", "231", "312", "321"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var perms []string
			permuations("", test.Input, &perms)
			assert.Equal(t, test.Expected, perms)
		})
	}
}

func TestHasAnagram(t *testing.T) {
	type TestCase struct {
		Input    []string
		Expected bool
	}

	tests := map[string]TestCase{
		"No anagrams": {
			Input:    []string{"1", "20", "3", "4", "5"},
			Expected: false,
		},
		"has an anagram": {
			Input:    []string{"100", "150", "215", "80", "152"},
			Expected: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Expected, hasAnagram(test.Input))
		})
	}
}

func TestGetChecksum(t *testing.T) {
	input := [][]string{
		{"1", "2", "3", "4", "5"},
		{"100", "150", "215", "80", "152"},
		{"500", "354", "50", "2", "99"},
		{"3001", "4", "1", "9", "500"},
	}
	expected := 3004

	assert.Equal(t, expected, getChecksum(input))
}
