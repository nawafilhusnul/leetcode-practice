package main

func containsDuplicate(nums []int) bool {
	checkMap := map[int]int{}
	for _, v := range nums {
		if _, ok := checkMap[v]; ok {
			return true
		}
		checkMap[v] = 0
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	xm := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := xm[v]; ok {
			return true
		}

		xm[v] = struct{}{}
	}

	return false
}
