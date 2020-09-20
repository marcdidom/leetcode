package problem720

import (
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	var ret string
	set := make(map[string]struct{}, len(words))
	for _, word := range words {
		if _, ok := set[word[:len(word)-1]]; !ok && len(word) != 1 {
			continue
		}
		set[word] = struct{}{}

		if len(word) > len(ret) {
			ret = word
		}
	}
	return ret
}
