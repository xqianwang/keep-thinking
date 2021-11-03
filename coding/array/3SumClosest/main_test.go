package _SumClosest

import "testing"

type Input struct {
	val []int
	target int
	expected int
}

func Test3SumClosest(t *testing.T)  {
	inputs := []Input{
		//{[]int{-1, 2 ,1, -4}, 1, 2},
		//{[]int{-1, -1, 1, 2, 6, -19, 25, 14, 10}, 5, 5},
		//{[]int{1, 1, 1, 1}, 2, 3},
		//{[]int{-1,0,1,1,55}, 3, 2},
		{[]int{-16,-2,17,-16,3,-7,-13,20,-4,12,5,13,-7,0,5,4,-17,-16,9,1,8,-6,0,-8,16,10,-6,9,-4,12,16,5,19,2,-9,-17,-8,-17,7,-10,2,20,-18,-20,-14,-6,6,17,-10,-8,18,-15,7,-9,13,13,-13,3,18,10,12,16,-6,-19,-6,-13,8,-5,16,5,8,-3,-9,-9,-5,14,14,-13,-18,13,15,-3,9,14,18,-14,-14,1,20,-4,-6,0,3,15,20,20,9,13,-8,-1,-2,6}, -58, -57},
	}
	for _, input := range inputs {
		result := _3SumClosest(input.val, input.target)
		if result != input.expected {
			t.Fatalf("Failed testing, expected %v actual is %v", input.expected, result)
		}
	}
}
