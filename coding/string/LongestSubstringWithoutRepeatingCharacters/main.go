package LongestSubstringWithoutRepeatingCharacters

//method 1, bad performance
func lengthOfLongestSubstring1(s string) int {
	result := make(map[rune]int)
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}
	for i := 0; i < len(s) - 1; i++ {
		temp := make(map[rune]int)
		temp[rune(s[i])] = 1
		for j := i + 1; j < len(s); j++ {
			if _, ok := temp[rune(s[j])]; !ok {
				temp[rune(s[j])] = 1
			} else {
				break
			}
		}
		if len(result) < len(temp) {
			result = temp
		}
	}
	return len(result)
}

//method 2  a little bit better
//Runtime:244 ms, faster than 16.22% of Go online submissions.
//Memory Usage:7.1 MB, less than 9.85% of Go online submissions.
func lengthOfLongestSubstring2(s string) int {
	result := make(map[rune]int)
	if len(s) <= 1 {
		return len(s)
	}
	for i := 0; i < len(s) - 1; i++ {
		temp := make(map[rune]int)
		temp[rune(s[i])] = i
		for j := i + 1; j < len(s); j++ {
			if _, ok := temp[rune(s[j])]; !ok {
				temp[rune(s[j])] = j
			} else {
				i = temp[rune(s[j])]
				break
			}
		}
		if len(result) < len(temp) {
			result = temp
		}
	}
	return len(result)
}

//method 3 best so far.... didn't think of it
//Runtime:4 ms, faster than 91.27% of Go online submissions.
//Memory Usage:2.5 MB, less than 93.09% of Go online submissions.
func lengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}