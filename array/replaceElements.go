package main

// replaceElements is good for memory
func replaceElements(arr []int) []int {
	res := make([]int, len(arr))

	index := 0
	rightIndex := index + 1
	maxRight := -1
	for index <= len(arr) {
		if index == len(arr)-1 {
			res[index] = -1
			break
		}

		if arr[rightIndex] > maxRight {
			maxRight = arr[rightIndex]
		}

		if rightIndex >= len(arr)-1 {
			res[index] = maxRight
			maxRight = -1
			index++
			rightIndex = index + 1
			continue
		}

		rightIndex++
	}

	return res
}

func findMaxElement(rightMax int, lastElement int) int {
	if lastElement > rightMax {
		rightMax = lastElement
	}
	return rightMax
}

// replaceElements2 is good for runtime
func replaceElements2(arr []int) []int {
	var rightMax int = -1
	for i := len(arr) - 1; i >= 0; i-- {
		var newMax int = findMaxElement(rightMax, arr[i])
		arr[i] = rightMax
		rightMax = newMax
	}
	return arr
}
