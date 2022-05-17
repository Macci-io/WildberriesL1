package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map

func main() {
	mutx := &sync.Mutex{}
	mapa := make(map[int]int)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func(mapa *map[int]int, i int, mutx *sync.Mutex, wg *sync.WaitGroup) {
			mutx.Lock()
			(*mapa)[i] = i * 2
			mutx.Unlock()
			wg.Done()
		}(&mapa, i, mutx, wg)
	}
	wg.Wait()
	for k, i := range mapa {
		time.Sleep(2 * time.Second)
		fmt.Println(k, i)
	}
}
