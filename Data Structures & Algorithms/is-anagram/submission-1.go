func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    countS, countT := make(map[rune]int), make(map[rune]int)
    for _, r := range s {countS[r]++}
    for _, r := range t {countT[r]++}

    for k, v := range countS {
        if countT[k] != v {
            return false
        }
    }
    return true
}
