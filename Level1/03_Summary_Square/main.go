package main

//-Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
import (
	"fmt"
	"sync"
)

func main() {
	pull := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 1)
	ch <- 0
	wg := sync.WaitGroup{}
	wg.Add(5)
	for _, i := range pull {
		go func(i int, ch chan int, wg *sync.WaitGroup) {
			ch <- (i * i) + <-ch
			wg.Done()
		}(i, ch, &wg)
	}
	wg.Wait()
	close(ch)
	fmt.Println(<-ch)
}
