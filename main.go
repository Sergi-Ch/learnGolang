package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//инициализация структуры с дефолтным значением (первый вариант)
	var John Person
	fmt.Printf("default first>> %T , %#v , \n", John, John)

	//инициализация структуры с дефолтным значением (второй вариант)
	John2 := Person{}
	fmt.Printf("default second>> %T , %#v , \n", John2, John2)

	//изменение полей структуры
	John.Name = "John"
	John.Age = 25
	fmt.Println("заполненные поля у Джона", John)

	//инициализация объекта и его полей

	Brad := Person{
		Name: "Brad",
		Age:  43,
	}
	fmt.Println(Brad)

	//инициализация сразу в одну строку по порядку полей
	Volodya := Person{"Volodya", 24}
	fmt.Println(Volodya)

	pVolodya := &Volodya         // получение указателя на объект Volodya
	fmt.Println((*pVolodya).Age) // сначала разыменовываем, чтобы потом обратиться к полю структуры
	fmt.Println(pVolodya.Age)    //синтаксический сахар go, обычно используют такой метод

	pIvan := &Person{"Ivan", 10} // создание сразу указателя на объект
	fmt.Println(pIvan, *pIvan)

	//создание анонимных структур, то есть без имени

	unnamedStruct := struct {
		Name, LastName, BirdDate string
	}{
		Name:     "NoName",
		LastName: "NoLastName",
		BirdDate: fmt.Sprintf("%s", time.Now()), //библиотека для работы со временем
		//или BirdDate: time.Now().String() будет одинаковый результат
	}
	fmt.Println(unnamedStruct)
}
