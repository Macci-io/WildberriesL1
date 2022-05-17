package main

import "fmt"

//type Methods interface {
//}

func (h *Human) walk(do string) { //-метод родительской структуры
	fmt.Printf("%s walking\n", do)
}

func (h *Human) somesault(do string) { //-метод родительской структуры
	fmt.Printf("%s somesaults\n", do)
}

type Human struct { //-родительская структура
	name string
	age  int
}

type Action struct { //-дочерняя структура
	Human //-встраиваем родительскую структуру
	do    string
}

func main() {
	action := Action{do: "I do"} //-объявляем и инициализируем переменную action, к которой присваиваем переменную do со значением
	action.walk(action.do)       //-приминяем метод walk унаследованный от родительского
	action.somesault(action.do)  //-приминяем метод somesault унаследованный от родительского

	human := Human{name: "Вася", age: 32}
	human.walk(human.name)
	human.somesault(human.name)

}
