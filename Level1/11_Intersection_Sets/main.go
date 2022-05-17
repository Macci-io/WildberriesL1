package main

import (
	"fmt"
)

//Реализовать пересечение двух неупорядоченных множеств;
func main() {
	Arr1 := []int{5, 7, 6, 3, 9, 8, 1, 4, 2}
	Arr2 := []int{10, 2, 5, 1, 15, 4, 11, 3}
	fmt.Println(Processing(Arr1, Arr2))
}

func Processing(arr1, arr2 []int) []int {
	mapa := make(map[int]int)
	res := make([]int, 0)

	for _, v := range arr1 {
		mapa[v] = 1
	}

	for _, v := range arr2 {
		if j, bl := mapa[v]; bl {
			if j == 1 {
				res = append(res, v)
				mapa[v]++
			}
		}
	}

	return res
}
