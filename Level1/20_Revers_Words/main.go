package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		var words string
		scan := bufio.NewScanner(os.Stdin)

		fmt.Println("Введите слова:")
		scan.Scan()
		words = scan.Text()
		arrWords := strings.Split(words, " ")
		left := 0
		right := len(arrWords) - 1
		for left < right {
			arrWords[left], arrWords[right] = arrWords[right], arrWords[left]
			left++
			right--
		}
		rev := strings.Join(arrWords, " ")
		fmt.Println(rev)
	}
}
