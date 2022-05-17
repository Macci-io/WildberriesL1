package main

import (
	"fmt"
)

func main() {
	var a int64
	var num int
	for {
		fmt.Println("Write a num of byte:")
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}
		if num < 0 || num > 64 {
			fmt.Println("Введите корректное значение!")
			continue
		}
		a ^= 1 << (num - 1)

		for i := 0; i < 64; i++ {
			if (a & (1 << i)) == 0 {
				print("0")
			} else {
				print("1")
			}
		}
		print("\n")
	}
}
