package main

import "fmt"

func main() {
	pattern := "ab"
	s := "abab"
	fmt.Println(wordPatternMatch(pattern, s))
}

func wordPatternMatch(pattern string, s string) bool {
	if len(pattern) > len(s) {
		return false
	}

	cMap := make(map[byte]string)
	sMap := make(map[string]byte)

	return search(cMap, 0, sMap, 0, pattern, s)
}

func search(cMap map[byte]string, cInd int, sMap map[string]byte, sInd int, pattern, s string) bool {
	if cInd >= len(pattern) {
		return false
	}

	c := pattern[cInd]

	for i := sInd + 1; i <= len(s); i++ {
		word := s[sInd:i]
		_, cok := cMap[c]
		_, sok := sMap[word]
		if cok && sok {
			if !search(cMap, cInd + 1, sMap, i + 1, pattern, s) {
				continue
			} else {
				return true
			}
		}
		if !cok && !sok {
			cMap[c] = word
			sMap[word] = c
			if !search(cMap, cInd + 1, sMap, i + 1, pattern, s) {
				delete(cMap, c)
				delete(sMap, word)
				continue
			} else {
				return true
			}
		} else {
			return false
		}
	}

	if cInd == len(pattern) - 1 {
		return true
	}
	return false
}
