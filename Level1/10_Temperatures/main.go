package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
func main() {
	temper := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mainPool := make(map[float64][]float64)
	var union float64

	for _, temp := range temper {
		union = math.Trunc(temp/10) * 10
		mainPool[union] = append(mainPool[union], temp)
	}
	for union, temp := range mainPool {
		fmt.Printf("%v : {", union)
		for _, v := range temp {
			fmt.Printf("%.1f, ", v)
		}
		fmt.Print("...}")
		fmt.Println("")
	}
}
