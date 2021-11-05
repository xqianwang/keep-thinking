package LongestSubstringWithoutRepeatingCharacters

import "testing"

type test struct {
	input string
	expected int
}

func TestLongestSubstring1(t *testing.T)  {
	tests := []test{
		{"abcabcbb", 3},
		{" ", 1},
		{"ch", 2},
		{"asd", 3},
		{"bbbbb", 1},
	}

	for _, te := range tests {
		r := lengthOfLongestSubstring1(te.input)
		if r != te.expected {
			t.Fatalf("Failed test, expected %d, actual %d", te.expected, r)
		}
	}
}

func TestLongestSubstring2(t *testing.T)  {
	tests := []test{
		{"abcabcbb", 3},
		{" ", 1},
		{"ch", 2},
		{"asd", 3},
		{"bbbbb", 1},
		{"aab", 2},
	}

	for _, te := range tests {
		r := lengthOfLongestSubstring2(te.input)
		if r != te.expected {
			t.Fatalf("Failed test, expected %d, actual %d", te.expected, r)
		}
	}
}

func TestLongestSubstring3(t *testing.T)  {
	tests := []test{
		{"abcabcbb", 3},
		{" ", 1},
		{"ch", 2},
		{"asd", 3},
		{"bbbbb", 1},
		{"aab", 2},
	}

	for _, te := range tests {
		r := lengthOfLongestSubstring3(te.input)
		if r != te.expected {
			t.Fatalf("Failed test, expected %d, actual %d", te.expected, r)
		}
	}
}