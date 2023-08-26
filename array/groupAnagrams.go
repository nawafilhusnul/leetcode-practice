package main

func groupAnagrams(strs []string) [][]string {
	checkArray := make(map[[26]int][]string)

	for _, str := range strs {
		var temp [26]int
		for _, s := range str {
			temp[s-'a']++
		}

		checkArray[temp] = append(checkArray[temp], str)
	}

	// checkArray =
	// {
	//     {1,0,0,0,1,0,0,0,0,0,0,...,1,0,0,0,...0}:[]{"eat","tea","ate"}
	// }

	var res [][]string
	for _, v := range checkArray {
		res = append(res, v)
	}

	return res
}
