package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество
func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	mapa := make(map[string]struct{})
	for _, v := range arr {
		mapa[v] = struct{}{}
	}

	result := make([]string, 0, len(mapa))

	for k, _ := range mapa {
		result = append(result, k)
	}
	fmt.Println(result)
}
