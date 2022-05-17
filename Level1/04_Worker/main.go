package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
func main() {
	ch := make(chan int)

	ctr, cancel := context.WithCancel(context.Background()) //*
	defer cancel()
	wg := new(sync.WaitGroup)
	exitc := make(chan os.Signal, 1)
	signal.Notify(exitc, os.Interrupt)

	go func() {
		select {
		case <-exitc:
			cancel()
		}
	}()

	var num int
	fmt.Print("Write a number:")

	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}
	wg.Add(num)
	for i := 1; num >= i; i++ {
		go worker(i, ch, wg, ctr)
	}
	go func() {
		for {
			ch <- rand.Intn(100)
		}
	}()
	wg.Wait()
}

func worker(i int, ch <-chan int, wg *sync.WaitGroup, ctr context.Context) {

	for {
		select {
		case data := <-ch:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Id: %d, data: %d\n", i, data)
		case <-ctr.Done():
			fmt.Println("Canceled with ctrl-c")
			wg.Done()
			return
		}
	}

	//fmt.Printf("Id: %d, data: %d\n", i, data)
	//time.Sleep(500 * time.Millisecond)

}
