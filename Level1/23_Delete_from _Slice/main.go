package main

import "fmt"

//Удалить i-ый элемент из слайса.
func main() {
	pool := make([]int, 10)
	for i := range pool {
		pool[i] = i
	}
	fmt.Println(Delete(pool, 5))
}

func Delete(pool []int, index int) []int {
	return append(pool[:index], pool[index+1:]...)
}
