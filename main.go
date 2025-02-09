package main

import (
	"fmt"
	"time"
)

func main() {
	// Вызов функций для демонстрации различных аспектов работы с каналами
	// nilChannel()
	// unbufferedChannel()
	// bufferChannel()
	forRange()
}

func nilChannel() {
	// Объявление канала без инициализации (nil канал)
	var nilChannel chan int
	// Вывод длины и емкости канала (оба будут равны 0)
	fmt.Printf("len: %d cap: %d\n", len(nilChannel), cap(nilChannel))
	// Попытка записи в nil канал вызовет deadlock
	// nilChannel <- 1
}

func unbufferedChannel() {
	// Создание небуферизованного канала
	unbufferedChannel := make(chan int)
	// Вывод длины и емкости канала (длина 0, емкость 0)
	fmt.Printf("len: %d cap: %d\n", len(unbufferedChannel), cap(unbufferedChannel))

	// Запуск горутины для записи в канал
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		// Запись значения в канал
		chanForWriting <- 1
	}(unbufferedChannel)
	// Чтение значения из канала
	value := <-unbufferedChannel
	println(value)

	// Запуск горутины для чтения из канала
	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		// Чтение значения из канала
		val := <-chanForReading
		println(val)
	}(unbufferedChannel)
	// Запись значения в канал
	unbufferedChannel <- 2
	// Закрытие канала
	close(unbufferedChannel)
}

func bufferChannel() {
	// Создание буферизованного канала с емкостью 2
	bufChannel := make(chan int, 2)
	// Вывод длины и емкости канала (длина 0, емкость 2)
	fmt.Printf("len: %d cap: %d\n", len(bufChannel), cap(bufChannel))
	// Запись значений в канал
	bufChannel <- 1
	bufChannel <- 2
	// Вывод длины и емкости канала (длина 2, емкость 2)
	fmt.Printf("len: %d cap: %d\n", len(bufChannel), cap(bufChannel))

	// Чтение значений из канала
	fmt.Println(<-bufChannel)
	fmt.Println(<-bufChannel)
	// После чтения значений можно снова записывать в канал
}

func forRange() {
	// Создание буферизованного канала с емкостью 3
	bufferedChannel := make(chan int, 3)

	// Слайс чисел для записи в канал
	number := []int{5, 6, 7, 8}

	// Запуск горутины для записи чисел в канал
	go func() {
		for _, num := range number {
			bufferedChannel <- num
		}
		// Закрытие канала после записи всех значений
		close(bufferedChannel)
	}()

	// Чтение значений из канала с проверкой на закрытие канала
	for {
		v, ok := <-bufferedChannel
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}

	// Повторное создание буферизованного канала
	bufferedChannel = make(chan int, 3)
	// Запуск горутины для записи чисел в канал
	go func() {
		for _, num := range number {
			bufferedChannel <- num
		}
		// Закрытие канала после записи всех значений
		close(bufferedChannel)
	}()
	// Чтение значений из канала с использованием range
	for v := range bufferedChannel {
		fmt.Printf("buffered: %#v \n", v)
	}

	// Создание небуферизованного канала
	unbufferedChannel := make(chan int)
	// Запуск горутины для записи чисел в канал
	go func() {
		for _, num := range number {
			unbufferedChannel <- num
		}
		// Закрытие канала после записи всех значений
		close(unbufferedChannel)
	}()
	// Чтение значений из канала с использованием range
	for v := range unbufferedChannel {
		fmt.Printf("unbuffered: %#v \n", v)
	}
}
