package main

import (
	"bufio"
	. "fmt"
	"os"
	"time"
)

//Реализовать собственную функцию sleep
func main() {
	Print("Insert value:")
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	Sleep(3 * time.Second)
	Println("Processing...")
	Sleep(3 * time.Second)
	Println("Processing...")
	Sleep(4 * time.Second)
	Print("Your value:")
	Sleep(4 * time.Second)
	Print(buf.Text())
}

func Sleep(tm time.Duration) {
	<-time.After(tm)
}
