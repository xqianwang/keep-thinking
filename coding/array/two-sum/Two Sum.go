package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		theOther := target - nums[i]
		if index, ok := m[theOther]; ok {
			return []int{i, index}
		}
		m[nums[i]] = i
	}
	return nil
}
