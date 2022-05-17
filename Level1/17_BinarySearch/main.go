package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка
func main() {

	pool := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println("Наш массив:", pool)
	toFind := 1
	index := Bsearch(pool, toFind)
	fmt.Printf("Индекс числа %d: %d", toFind, index)
}

func Bsearch(pool []int, num int) int {
	left := 0
	right := len(pool) - 1

	for left <= right {
		mid := (left + right) / 2
		if pool[mid] == num {
			return mid
		} else if num > pool[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
