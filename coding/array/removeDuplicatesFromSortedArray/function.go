package removeDuplicatesFromSortedArray

type flag struct {
	val int
	dup bool
	pos int
}

func removeDuplicates(nums []int) int {
	last, finder, length := 0, 0, len(nums)
	for last < length -1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == length - 1 {
				return last + 1
			}
		}
		nums[last] = nums[finder]
		last++
	}
	return last + 1
}
