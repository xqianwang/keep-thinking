package palindrome_number

import "testing"

func TestFunction(t *testing.T)  {
	test1 := 1991
	result1 := isPalindrome(test1)
	if !result1 {
		t.Fatalf("input is %d, expected value %v, actual value %v", test1, true, result1)
	}

	test2 := 0
	result2 := isPalindrome(test2)
	if !result2 {
		t.Fatalf("input is %d, expected value %v, actual value %v", test2, true, result2)
	}

	result3 := isPalindrome2(test1)
	if !result3 {
		t.Fatalf("input is %d, expected value %v, actual value %v", test1, true, result3)
	}

	result4 := isPalindrome2(test2)
	if !result4 {
		t.Fatalf("input is %d, expected value %v, actual value %v", test2, true, result4)
	}
}
