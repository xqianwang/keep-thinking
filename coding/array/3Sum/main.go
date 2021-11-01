package _Sum

import (
	"reflect"
	"sort"
)

//time complexity too high. Failed solution
func _3Sum1(nums []int) [][]int {
	//inputMap := make(map[int]Com)
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums) - 2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i - 1]) {
			low := i + 1
			high := len(nums) - 1
			target := 0 - nums[i]
			for j := low; j < high; j++ {
				k := high
				for j < k {
					exist := false
					if nums[k] + nums[j] == target {
						for _, v := range result {
							if reflect.DeepEqual(v[1:], []int{nums[j], nums[k]}) {
								exist = true
							}
						}
						if exist == false {
							result = append(result, []int{nums[i], nums[j], nums[k]})
						}
					}
					k--
				}
			}
		}
	}
	return result
}
func _3Sum2(nums []int) [][]int {
	//inputMap := make(map[int]Com)
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums) - 2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i - 1]) {
			low := i + 1
			high := len(nums) - 1
			target := 0 - nums[i]
			for low < high {
				if nums[low] + nums[high] == target {
					result = append(result, []int{nums[i], nums[low], nums[high]})
					for low < high && nums[low] == nums[low + 1]  {
						low++
					}
					for low < high && nums[high] == nums[high - 1] {
						high --
					}
					low++
					high--
				} else if nums[low] + nums[high] > target {
					high--
				} else {
					low++
				}
			}
		}
	}
	return result
}

func _3Sum3(nums []int) [][]int  {
	result, start, index, end, length := make([][]int, 0), 0, 0, 0, len(nums)
	sort.Ints(nums)
	for index = 1; index < length - 1; index++ {
		start, end = 0, length - 1
		if index > 1 && nums[index] == nums[index - 1] {
			start = index - 1
		}
		for start < index && end > index {
			if nums[start]+nums[index]+nums[end] == 0 {
				if start > 0 && nums[start] == nums[start-1] {
					start++
					continue
				}
				if end < length-1 && nums[end] == nums[end+1] {
					end--
					continue
				}
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if nums[start] + nums[index] + nums[end] > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}