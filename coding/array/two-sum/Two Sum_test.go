package two_sum

import (
	"fmt"
	"testing"
)

type params struct {
	num []int
	target int
}

type expected struct {
	expected []int
}

type test_instance struct {
	params
	expected
}

func TestTwoSum(t *testing.T) {
	ti := []test_instance {
		{
			params{[]int{1, 2, 3, 9, 10}, 12},
			expected{[]int{2, 3}},
		},
		{
			params{[]int{0, 2, 15, 32}, 47},
			expected{[]int{2, 3}},
		},
		{
			params{[]int{-3, 24, -9, 88, 108, -8}, 100},
			expected{[]int{4, 5}},
		},
	}
	for _, instance := range ti{
		_, p := instance.expected, instance.params

		fmt.Printf("[input]:%v  [output]: %v\n", p, twoSum(p.num, p.target))
	}
}
