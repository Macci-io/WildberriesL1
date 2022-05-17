package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	var sec int
	fmt.Print("Write a second:")
	_, err := fmt.Scan(&sec)
	if err != nil {
		fmt.Println(err)
		return
	}
	timend, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(sec))
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go sender(ch)
	go reader(ch, wg, timend)
	wg.Wait()
}

func sender(ch chan int) {
	for {
		select {
		default:
			ch <- rand.Intn(100)
		}
	}
}

func reader(ch chan int, wg *sync.WaitGroup, timend context.Context) {
	for {
		select {
		case r := <-ch:
			//time.Sleep(500 * time.Millisecond)
			fmt.Println(r)
		case <-timend.Done():
			fmt.Println("Timeout")
			wg.Done()
			return
		}
	}
}
