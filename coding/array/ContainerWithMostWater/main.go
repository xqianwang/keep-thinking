package ContainerWithMostWater

func containerWithMostWater(input []int) int  {
	i, j := 0, len(input) - 1
	var area, max int
	for i < j {
		if input[i] <= input[j] {
			area = input[i] * (j - i)
			i++
		} else {
			area = input[j] * (j - i)
			j--
		}
		if area >= max {
			max = area
		}
	}
	return max
}
