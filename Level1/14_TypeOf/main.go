package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}
func main() {

	check(Method)
}

func Method(interf interface{}) {
	fmt.Println(reflect.TypeOf(interf))
}

func check(interf func(interface{})) {
	interf(5)
	interf("Hello")
	interf(true)
	interf(new(chan int))
}
