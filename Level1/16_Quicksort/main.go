package main

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
func main() {
	arr := []int{2, 5, 10, 4, 9, 8, 6, 7, 1}
	fmt.Println("До:", arr)
	QuickSort(arr)
	fmt.Println("После:", arr)
}

func QuickSort(arr []int) []int {

	left := 0
	right := len(arr) - 1

	if len(arr) < 2 {
		return arr
	}

	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return arr
}
