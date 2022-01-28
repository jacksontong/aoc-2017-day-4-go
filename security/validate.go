package security

import (
	"reflect"
	"sort"
	"strings"
)

// A passphrase consists of a series of words (lowercase letters) separated by spaces
// a valid passphrase must contain no duplicate words.
func ValidatePassword(pass string) bool {
	checker := make(map[string]bool)
	words := strings.Split(pass, " ")

	for i, word := range words {
		_, ok := checker[word]
		if !ok {
			checker[word] = true
			for j, next := range words {
				if i == j {
					continue
				}

				if isAnagram(word, next) {
					checker[word] = false
					// mark the next word as invalid
					// so it'll skip the inner loop
					checker[next] = false
				}
			}
		} else {
			// the word is duplicated
			checker[word] = false
		}
	}

	for _, ok := range checker {
		if !ok {
			return false
		}
	}

	return true
}

// Check if 2 given words are anagrams of each other
func isAnagram(w1, w2 string) bool {
	first := strings.Split(w1, "")
	second := strings.Split(w2, "")
	sort.Strings(first)
	sort.Strings(second)

	return reflect.DeepEqual(first, second)
}
