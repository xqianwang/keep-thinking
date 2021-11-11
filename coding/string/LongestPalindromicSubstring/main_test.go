package LongestPalindromicSubstring

import "testing"

type input struct {
	s string
	expectedS string
}
func TestLongestPalindrome(t *testing.T) {
	inputs := []input {
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"ccc", "ccc"},
		{"ccd", "cc"},
	}

	for _, i := range inputs {
		result := longestPalindrome(i.s)
		if i.expectedS != result {
			t.Fatalf("Wrong result, expected %s, actual value %s", i.expectedS, result)
		}
	}
}
