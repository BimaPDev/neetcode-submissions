func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range strs {
		charCount := make([]int, 26)
		for _, char := range word {
			charCount[char-'a']++
		}
		key := fmt.Sprint(charCount)
		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}