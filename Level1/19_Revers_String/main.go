package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
func main() {
	var str string
	fmt.Println("Введите слово: ")
	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}
		sym := []rune(str)
		left := 0
		right := len(sym) - 1
		for left < right {
			sym[left], sym[right] = sym[right], sym[left]
			left++
			right--
		}
		fmt.Print("Результат:")
		fmt.Println(string(sym))
	}
}
