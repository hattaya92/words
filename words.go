package words

func WordCount(s string) map[string]int {
	var word map[string]int
	word = make(map[string]int)
	k := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == ',' || s[i] == '.' {
			word[s[k:i]] += 1
			k = i
		}
	}
	return word
}
