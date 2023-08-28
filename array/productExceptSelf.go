package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	forward := 1
	for idx, num := range nums {
		res[idx] = forward
		forward *= num
	}
	fmt.Println(res)
	backward := 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		res[idx] *= backward
		backward *= nums[idx]
	}

	return res
}
