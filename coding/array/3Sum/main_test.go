package _Sum

import (
	"reflect"
	"testing"
)

type input struct {
	val []int
	expected [][]int
}

func Test3Sum(t *testing.T)  {
	suits := []input {
		{[]int{-1, 0, 1, -1, 2, -4}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		{[]int{0, 0, 0, 0}, [][]int{[]int{0, 0, 0}}},
	}

	for _, v := range suits {
		result := _3Sum1(v.val)
		if !reflect.DeepEqual(result, v.expected) {
			t.Fatalf("Failed test, expected %v, actual %v", v.expected, result)
		}
	}

	for _, v := range suits {
		result := _3Sum2(v.val)
		if !reflect.DeepEqual(result, v.expected) {
			t.Fatalf("Failed test, expected %v, actual %v", v.expected, result)
		}
	}

	for _, v := range suits {
		result := _3Sum3(v.val)
		if !reflect.DeepEqual(result, v.expected) {
			t.Fatalf("Failed test, expected %v, actual %v", v.expected, result)
		}
	}
}
