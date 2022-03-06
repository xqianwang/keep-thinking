package main

import (
	"fmt"
	"sort"
)

func main()  {
	test := "aabc"
	fmt.Println(minDeletions(test))
}
//"abcaabbc"
//332 // 322

func minDeletions(s string) int {
	ret := 0
	count := make([]int, 26)
	for _, r := range s {
		count[r - 'a']++
	}

	sort.Ints(count)
	var used []int
	for i := 25; i >= 0; i-- {
		if count[i] > 0 && contains(used, count[i]) {
			for count[i] > 0 && contains(used, count[i]) {
				count[i]--
				ret++
			}
			if count[i] > 0 {
				used = append(used, count[i])
			}
		} else {
			used = append(used, count[i])
		}
	}

	return ret
}

func contains(list []int, e int) bool {
	for _, v := range list {
		if v == e {
			return true
		}
	}

	return false
}