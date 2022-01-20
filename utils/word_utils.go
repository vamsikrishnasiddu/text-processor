package utils

import "strings"

// WordCount gives a map of the words with their
//occurence in the text
func WordCount(st string) map[string]int {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}
