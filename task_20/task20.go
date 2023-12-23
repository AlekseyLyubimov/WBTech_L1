package main

import "strings"

func main() {
	str := "Карл у Клары украл кораллы"
	println(flipper(str))
}

func flipper(str string) string {
	words := strings.Split(str, " ")
	for left, right := 0, len(words)-1; left < right; {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}
