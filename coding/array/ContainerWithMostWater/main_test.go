package ContainerWithMostWater

import (
	"testing"
)

type Cases struct {
	input []int
	expectedResult int
}

func TestContainerWithMistWater(t *testing.T)  {
	cases := []Cases {
		{[]int{1,8,6,2,5,4,8,3,7}, 49},
		{[]int{1,6,22,19999,15,47,9,1}, 94},
	}

	for _, i := range cases {
		result := containerWithMostWater(i.input)
		if result != i.expectedResult {
			t.Fatalf("Incorrect result, expect %d, actual value %d", i.expectedResult, result)
		}
	}
}