package main

/*Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/
func main() {
	pool := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go FirstChannel(pool, ch1)
	go SecondChannel(ch1, ch2)
	//time.Sleep(5 * time.Second)
	for v := range ch2 {
		println("Результат:", v)
	}
}

func FirstChannel(pool []int, ch1 chan int) {
	for _, x := range pool {
		ch1 <- x
	}
	close(ch1)
}

func SecondChannel(ch1 chan int, ch2 chan int) {
	for y := range ch1 {
		ch2 <- y * 2
	}
	close(ch2)
}
