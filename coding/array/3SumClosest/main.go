package _SumClosest

import (
	"math"
	"sort"
)

func _3SumClosest(nums []int, target int) int {
	sort.Ints(nums)
	low, index, high, length, closest := 0, 1, 0, len(nums), math.MaxInt32
	for index = 1; index < length - 1; index++ {
		low, high = 0, length - 1
		for low < index && high > index {
			sum := nums[low] + nums[index] + nums[high]
			if sum == target {
				return sum
			} else if sum > target {
				high--
			} else {
				low++
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			}
		}
	}
	return closest
}
