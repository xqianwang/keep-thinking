package main

import "fmt"

func main()  {
	fmt.Println(sumZero(5))
}

func sumZero(n int) []int {
	var result []int
	for i := 1; i <= n / 2; i++ {
		result = append(result, i)
		result = append(result, 0 - i)
	}
	if n % 2 != 0 {
		result = append(result, 0)
	}

	return result
}
