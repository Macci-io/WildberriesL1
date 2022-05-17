package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false

*/
func main() {
	var str string
	for {
		fmt.Print("Insert value: ")
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}
		fmt.Println(str, "-", check(str))
	}
}

func check(str string) bool {
	sym := []rune(strings.ToLower(str))
	for i := 0; i < len(sym); i++ {
		for j := i + 1; j < len(sym); j++ {
			if sym[i] == sym[j] {
				return false
			}
		}
	}
	return true
}
