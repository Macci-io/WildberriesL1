package main

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

import (
	"fmt"
	"sync"
)

type CountSet struct {
	sync.Mutex
	count int
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1000000)
	obj := CountSet{count: 0}

	for i := 0; i < 1000000; i++ {
		go Concurrency(wg, &obj)
	}
	wg.Wait()
	fmt.Println(obj.count)
}

func Concurrency(wg *sync.WaitGroup, obj *CountSet) {
	//obj.Lock()
	obj.count++
	//obj.Unlock()
	wg.Done()
}
