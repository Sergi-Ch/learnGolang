package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Вызываем функции для демонстрации работы с выводом и вводом
	outputDemo()
	inputDemo()
}

// outputDemo демонстрирует различные способы вывода данных
func outputDemo() {
	// 1. Вывод в консоль с помощью fmt.Println
	// Возвращает количество записанных байт и ошибку
	count, err := fmt.Println("some text")
	if err != nil {
		log.Fatalf("Ошибка при выводе: %v", err)
	}
	fmt.Printf("Количество записанных байт: %d\n", count)

	// 2. Форматированный вывод в стандартный поток вывода (stdout)
	count, err = fmt.Fprintf(os.Stdout, "Stdout from FprintF\n")
	if err != nil {
		log.Fatalf("Ошибка при выводе в stdout: %v", err)
	}
	fmt.Printf("Количество записанных байт: %d\n", count)

	// 3. Форматированный вывод в стандартный поток ошибок (stderr)
	count, err = fmt.Fprintf(os.Stderr, "Stderr from FprintF\n")
	if err != nil {
		log.Fatalf("Ошибка при выводе в stderr: %v", err)
	}
	fmt.Printf("Количество записанных байт: %d\n", count)

	// 4. Создание файла для записи
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("Ошибка при создании файла: %v", err)
	}
	defer func() {
		// Закрытие файла при завершении функции
		err := file.Close()
		if err != nil {
			log.Fatalf("Ошибка при закрытии файла: %v", err)
		}
	}()

	// Запись в файл с помощью fmt.Fprintf
	message := "hello it's check for writing and reading,\nmy friend"
	count, err = fmt.Fprintf(file, message)
	if err != nil {
		log.Fatalf("Ошибка при записи в файл: %v\n", err)
	}
	fmt.Printf("В файл записано %d байт\n", count)
}

// inputDemo демонстрирует различные способы ввода данных
func inputDemo() {
	var (
		text1 string
		text2 string
	)

	// 1. Чтение из стандартного ввода (консоли)
	fmt.Println("Введите два слова через пробел:")
	count, err := fmt.Scan(&text1, &text2)
	if err != nil {
		log.Fatalf("Ошибка при чтении из консоли: %v", err)
	}
	fmt.Printf("Прочитано: %s %s (количество успешно прочитанных значений: %d)\n",
		text1, text2, count)

	// 2. Открытие файла для чтения
	file, err := os.Open("test.txt")
	if err != nil {
		log.Panicf("Ошибка при открытии файла: %v", err)
	}
	defer func() {
		// Закрытие файла при завершении функции
		err := file.Close()
		if err != nil {
			log.Fatalf("Ошибка при закрытии файла: %v", err)
		}
	}()

	// Чтение из файла с помощью fmt.Fscan
	// Fscan читает до пробелов или переноса строки (с припиской ln)
	count, err = fmt.Fscan(file, &text1, &text2)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
	}
	fmt.Printf("Из файла прочитано: %s %s (количество значений: %d)\n",
		text1, text2, count)
}
