func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)

    for _, words := range strs {
        var charCount [26]int
        for _, ch := range words {
            charCount[int(ch-'a')]++
        }
        key := fmt.Sprint(charCount)
        groups[key] = append(groups[key], words)
    }
    result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result

}
