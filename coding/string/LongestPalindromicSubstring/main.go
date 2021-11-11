package LongestPalindromicSubstring

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	if len(s) == 2  {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}

	var max string
	for i := 0; i < len(s) - 1; i++ {
		result := testPalindrome(s, i)
		if len(max) < len(result) {
			max = result
		}
	}
	return max
}

func testPalindrome(s string, i int) string  {
	result := ""
	if i == 0 {
		for j := i+1; j < len(s); j++ {
			if s[j] != s[i] {
				result = s[:j]
				return result
			}
		}
	} else {
		j, k := i - 1, i + 1
		for j >= 0 && k <= len(s) - 1 {
			if s[j] != s[k] {
				result = s[j+1:k]
				break
			}
			result = s[j:k+1]
			j--
			k++
		}
		j = i + 1
		for i >= 0 && j <= len(s) - 1 {
			if s[i] != s[j] {
				if j - i > 1 {
					if result == "" || len(s[i:j]) > len(result) {
						result = s[i+1:j]
					}
				} else {
					if result == "" {
						result = string(s[i])
					}
				}
				break
			}
			if len(s[i:j + 1]) > len(result)  {
				result = s[i:j + 1]
			}
			i--
			j++
		}
	}
	return result
}