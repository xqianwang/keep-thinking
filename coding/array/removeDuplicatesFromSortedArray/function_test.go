package removeDuplicatesFromSortedArray

import (
	"fmt"
	"testing"
)

type instance struct {
	list []int
	answer int
}

func TestFunction(t *testing.T) {
	instances := []instance {
		{[]int{0, 0, 1, 2, 2}, 3},
		{
			[]int{1, 1, 2, 4, 5, 6},
			5,
		},
		{
			[]int{100, 101, 101, 101, 103, 145},
			4,
		},
	}

	for _, v := range instances {
		fmt.Printf("input is %v, the actual value is: %v", v, removeDuplicates(v.list))
	}
}
