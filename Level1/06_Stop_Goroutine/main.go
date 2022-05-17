package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func main() {
	CancelCtx(5)
	TimeOutCtx(5)
	ChannelExit(5)
	DeadlineCtx(5)
	OsExit()
}

func CancelCtx(num int) {
	time.Sleep(1 * time.Second)
	fmt.Println("With cancel context is started:")
	time.Sleep(1 * time.Second)
	wg := new(sync.WaitGroup)
	ctx, cancl := context.WithCancel(context.Background())
	wg.Add(num)
	for i := 1; i <= num; i++ {
		time.Sleep(time.Second)
		go ThreadWithContext(ctx, wg, i)
	}
	<-time.After(time.Second)
	fmt.Println("...time is ticking...")
	cancl()
	wg.Wait()
	fmt.Println("Goroutine stop working")
	fmt.Println("")
}

func ThreadWithContext(ctx context.Context, wg *sync.WaitGroup, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\tCancel %d\n", id)
			wg.Done()
			return
		default:
			fmt.Printf("\tProcessing... %d\n", id)
			time.Sleep(10 * time.Second)
		}
	}
}

func TimeOutCtx(num int) {
	time.Sleep(2 * time.Second)
	fmt.Println("With timout context is started:")
	wg := &sync.WaitGroup{}
	wg.Add(num)
	ctx, cancl := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancl()
	for i := 1; i <= num; i++ {
		go ThreadWithContext(ctx, wg, i)
		time.Sleep(time.Second)
	}
	fmt.Println("...time is ticking...")
	wg.Wait()
	fmt.Println("Goroutine stop working")
	fmt.Println("")
}
func ChRout(ch <-chan struct{}, i int, wg *sync.WaitGroup) {
	<-ch
	fmt.Println("Done channel", i)
	wg.Done()
}
func ChannelExit(num int) {
	time.Sleep(2 * time.Second)
	fmt.Println("With channel is started")
	time.Sleep(2 * time.Second)
	wg := new(sync.WaitGroup)
	wg.Add(num)
	ch := make(chan struct{})
	for i := 1; i <= num; i++ {
		time.Sleep(time.Second)
		go ChRout(ch, i, wg)
	}
	//time.Sleep(2 * time.Second)
	close(ch)
	fmt.Println("Close channel")
	wg.Wait()
	fmt.Println("")
}

func DeadlineCtx(num int) {
	time.Sleep(2 * time.Second)
	fmt.Println("With Deadline context is started:")
	wg := new(sync.WaitGroup)
	wg.Add(num)
	ctx, cancl := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancl()
	for i := 1; i <= num; i++ {
		go ThreadWithContext(ctx, wg, i)
		time.Sleep(time.Second)
	}
	wg.Wait()
	fmt.Println("Deadline context is stopping")
	fmt.Println("")
}

func OsExit() {
	time.Sleep(2 * time.Second)
	fmt.Println("")
	fmt.Println("Exit with code: -1")
	os.Exit(-1)
}
