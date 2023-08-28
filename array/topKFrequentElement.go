package main

func topKFrequent(nums []int, k int) []int {
	countEach := make(map[int]int)
	for _, num := range nums {
		countEach[num]++
	}

	counts := make([][]int, len(nums)+1)
	for num, count := range countEach {
		counts[count] = append(counts[count], num)
	}

	res := []int{}
	for i := len(counts) - 1; i >= 0; i-- {
		res = append(res, counts[i]...)
		if len(res) == k {
			return res
		}
	}
	return res
}
