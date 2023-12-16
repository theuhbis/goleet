package anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	memo := make(map[byte]int, 26)

	for i := 0; i < len(s); i++ {
		if _, ok := memo[s[i]]; ok {
			memo[s[i]]++
		} else {
			memo[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := memo[t[i]]; ok {
			memo[t[i]]--
			if memo[t[i]] == 0 {
				delete(memo, t[i])
			}

		} else {
			return false
		}
	}
	
	return len(memo) == 0
}