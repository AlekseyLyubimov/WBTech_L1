package main


func main() {
	str := "abcdefghij"
	println(flipper(str))
}

func flipper(str string) string {
	runes := []rune(str)

	for left, right := 0, len(runes) - 1; left < right; {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
