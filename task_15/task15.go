package main

import "strings"

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
/*
	няп проблема в том, что ХугеСтринг целиком останется в памяти, потому что слайс на неё ссылается, и вместо 100 килобайт будет занят 1 кибибайт
	нужно сделать глубокую копию через strings.Clone(v[:100])
*/

var justString string

func someFunc_corrected() {
  v := createHugeString(1 << 10)
  justString = strings.Clone(v[:100])
}

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc_corrected()
  println(justString)
}

func createHugeString(i int) (result string) {
	for i > 0 {
		result += "W"
		i--
	}
  return result
}
