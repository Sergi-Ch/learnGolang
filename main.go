package main

import "fmt"

func main() {
	//default value
	var intPointer *int // раз не даем значение - он сразу присваивает nill
	fmt.Printf("defuault value %T %#v \n", intPointer, intPointer)

	//получение не nill указателей (более безопасно)
	//указатель на существующую переменную
	var a int64 = 10 // сама переменная
	fmt.Printf("%T %#v \n", a, a)

	var pointerA *int64 = &a // присваиваем переменной указатель на а, тип можно указывать, можно нет, он сам понимает
	fmt.Printf("указатель на существующую переменную тип %T значение адреса %#v значение с разыменовыванием %#v \n", pointerA, pointerA, *pointerA)

	// указатель с помощью new (безопаснее, чем самый 1 вариант из-за дефолтного значения)
	var newPointer = new(float64)
	fmt.Printf("Указатель с помощью new %T %#v вниманеие на значение по умолчанию >> %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 7
	fmt.Printf("Указатель с помощью new %T %#v вниманеие на значение >> %#v \n", newPointer, newPointer, *newPointer)

}
