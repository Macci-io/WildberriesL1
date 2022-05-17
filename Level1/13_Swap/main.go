package main

import "fmt"

func main() {
	var a, b int64 = 123, 456
	fmt.Printf("before: a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("after:  a = %d, b = %d\n", a, b)
}
