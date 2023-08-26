package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	checkMap := map[byte]int{}
	for _, v := range s {
		checkMap[byte(v)]++
	}

	for _, v := range t {
		if val, ok := checkMap[byte(v)]; ok {
			if val <= 0 {
				return false
			}
			checkMap[byte(v)]--
		} else {
			return false
		}
	}

	return true

}

// isAnagram2 is better at runtime and memory
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for idx := 0; idx < len(s); idx++ {
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
	}

	for idx := 0; idx < len(freq); idx++ {
		if freq[idx] != 0 {
			return false
		}
	}

	return true
}
