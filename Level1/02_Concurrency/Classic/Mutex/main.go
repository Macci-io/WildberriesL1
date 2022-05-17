package main

import (
	"fmt"
	"sync"
)

var (
	mt sync.Mutex //-
)

/*
Mutex
*/
func main() {
	wg := new(sync.WaitGroup) //-
	/*

	 */
	pull := [5]int{2, 4, 6, 8, 10} //-

	wg.Add(len(pull))        //-
	for _, i := range pull { //-

		go calc(i, wg) //-

	}
	wg.Wait()         //-
	fmt.Println(pull) //-
}

func calc(i int, group *sync.WaitGroup) {
	mt.Lock()          //-
	defer mt.Unlock()  //-
	defer group.Done() //-
	fmt.Println(i * i) //-
}
