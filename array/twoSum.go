package main

func twoSum(nums []int, target int) []int {
	hMap := make(map[int]int)

	// might use for range also
	for idx := 0; idx < len(nums); idx++ {
		expectedNumber := target - nums[idx]
		if val, ok := hMap[expectedNumber]; ok {
			return []int{val, idx}
		}
		hMap[nums[idx]] = idx
	}

	return nil
}
