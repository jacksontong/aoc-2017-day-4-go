package security

import "strings"

// A passphrase consists of a series of words (lowercase letters) separated by spaces
// a valid passphrase must contain no duplicate words.
func ValidatePassword(pass string) bool {
	checker := make(map[string]bool)

	for _, word := range strings.Split(pass, " ") {
		_, ok := checker[word]
		if !ok {
			checker[word] = true
		} else {
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
