package palindrome_number

import (
	"strconv"
)

//bad performance solution
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else if x % 10 == 0 {
		return false
	} else {
		arr := make([]int, 0, 32)
		for x > 0 {
			arr = append(arr, x % 10)
			x = x/10
		}
		size := len(arr)
		i, j := 0, size - 1
		for i < j {
			if arr[i] != arr[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
}

//method 2
func isPalindrome2(x int) bool {
	if x == 0 {
		return true
	}
	xstring := strconv.Itoa(x)
	i, j := 0, len(xstring) - 1
	for i < j {
		if xstring[i] != xstring[j] {
			return  false
		}
		i++
		j--
	}
	return true
}