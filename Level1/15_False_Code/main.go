package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/
func main() {
	var str string

	v := createHugeString(1 << 10)

	str = strings.Clone(v[:100])
	fmt.Printf(str)
}

func createHugeString(size int) string {
	return strings.Repeat("-|_|-", size)
}
