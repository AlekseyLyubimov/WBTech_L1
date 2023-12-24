package main

import "strings"

func main() {
	println(check_unique("abcd"))
	println(check_unique("abCdefAaf"))
	println(check_unique("aabcd"))
}

func check_unique(str string) bool {
	m := make(map[rune]bool)

	for _, ch := range strings.ToUpper(str) {
		_, ok := m[ch]
		if ok {
			return false
		} else {
			m[ch] = true
		}
	}

	return true
}