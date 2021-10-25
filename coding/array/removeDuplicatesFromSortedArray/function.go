package removeDuplicatesFromSortedArray

type flag struct {
	val int
	dup bool
	pos int
}
//solution1
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	last, finder, length := 0, 0, len(nums)
	for last < length -1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == length {
				return last + 1
			}
		}
		last++
		nums[last] = nums[finder]

	}
	return last + 1
}

//solution2
func removeDuplicates2(nums []int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[p] = nums[i]
			p++
		}
	}
	return p
}
