package main

import (
	"fmt"
)

func main() {
	// Запуск горутины, которая выводит числа от 0 до 99
	// go showNubmers(100)

	// Передача управления другой горутине
	// runtime.Gosched()

	// Установка максимального количества процессоров, используемых для выполнения горутин
	// runtime.GOMAXPROCS(1)

	// Вывод количества доступных процессоров
	// fmt.Println(runtime.NumCPU())

	// Регистрация defer функций, которые будут выполнены перед завершением функции main
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println("exit")

	// Вызов функции sum и вывод результата
	// println(sum(1, 2))

	// Вызов функции deferValues для демонстрации работы defer
	// deferValues()

	// Вызов функции makePanic для демонстрации работы panic и recover
	makePanic()

	// Этот код не будет выполнен, так как makePanic вызывает панику
	fmt.Println("after panic")
}

// Функция, которая выводит числа от 0 до numb-1
func showNubmers(numb int) {
	for i := 0; i < numb; i++ {
		fmt.Println(i)
	}
}

// Функция, которая суммирует два числа и умножает результат на 2 с помощью defer
func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()
	sum = x + y
	return
}

// Функция, демонстрирующая различные способы использования defer
func deferValues() {
	// Правильный вариант: defer регистрирует вывод значения i на каждой итерации
	for i := 0; i < 10; i++ {
		defer fmt.Println("first", i)
	}

	// Неправильный вариант: defer регистрирует вывод значения i, которое изменяется на каждой итерации
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("second", i)
		}()
	}

	// Исправленный вариант: defer регистрирует вывод значения k, которое является копией i
	for i := 0; i < 10; i++ {
		k := i
		defer func() {
			fmt.Println("first fix of second", k)
		}()
	}

	// Исправленный вариант: defer регистрирует вывод значения k, которое передается как аргумент
	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println("second fix of second", k)
		}(i)
	}
}

// Функция, демонстрирующая работу panic и recover
func makePanic() {
	defer func() {
		// Восстановление после паники и вывод значения, переданного в panic
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	// Вызов паники с сообщением "some panic"
	panic("some panic")
}
