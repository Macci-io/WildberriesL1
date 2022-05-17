package main

import (
	"fmt"
)

//-конкуренция на примере канала
func main() {
	//wg := new(sync.WaitGroup)
	pull := []int{2, 4, 6, 8, 10} //-
	message := make(chan int, 1)  //-

	//var wg sync.WaitGroup
	//wg.Add(5)
	for _, i := range pull { //-
		go calc(i, message) //-
	}
	for i := 0; i < len(pull); i++ { //-
		fmt.Println(<-message) //-
	}
	//close(message)
}

func calc(i int, msg chan int) {
	msg <- i * i //-
	//wg.Done()
}

/*

 */
