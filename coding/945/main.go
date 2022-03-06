package main

import (
	"fmt"
	"sort"
)

func main()  {
	test := []int{3,2,1,2,1,7}
	fmt.Println(minIncrementForUnique(test))
}

func minIncrementForUnique(nums []int) int {
	ret := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i - 1] {
			ret += nums[i - 1] - nums[i] + 1
			nums[i] = nums[i - 1] + 1
		}
	}

	return ret
}