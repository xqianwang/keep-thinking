package main

import "fmt"

func main()  {
	test := " test     haha hello    world   "
	fmt.Println(reverseWords(test))
}


func reverseWords(s string) string {
	left, right := 0, len(s) - 1
	chars := []rune{}
	for left <= right && s[left] == 32 {
		left++
	}
	for left <= right && s[right] == 32 {
		right--
	}

	s = s[left:right + 1]

	for _, c := range s {
		if c != ' ' {
			chars = append(chars, c)
		} else if chars[len(chars) - 1] != ' ' {
			chars = append(chars, c)
		}
	}

	left, right = 0, len(chars) - 1
	for left <= right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}

	lastIndex := 0
	for i, c := range chars {
		if c == ' ' {
			l, r := lastIndex, i-1
			for l <= r {
				chars[l], chars[r] = chars[r], chars[l]
				l++
				r--
			}
			lastIndex = i + 1
		}
	}

	l, r := lastIndex, len(chars) - 1
	for l <= r {
		chars[l], chars[r] = chars[r], chars[l]
		l++
		r--
	}

	return string(chars)
}