package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Раскомментируйте нужные функции для тестирования
	// CreateFile()
	// WriteTo()
	// readFrom()
	// ioCopy()
	// createFileBufio()
}

// CreateFile создает файл test.txt и записывает в него числа от 0 до 99
func CreateFile() {
	start := time.Now()

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// Всегда закрываем файл, даже если произошла ошибка при записи
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for i := 0; i < 100; i++ {
		// Используем более идиоматичный цикл for
		if _, err = file.WriteString(fmt.Sprintf("%d\n", i)); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("CreateFile duration:", time.Since(start))
}

// WriteTo читает содержимое файла test.txt и выводит его в stdout
func WriteTo() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err := file.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nWritten %d bytes\n", count)
}

// readFrom создает файл test2.txt и копирует в него содержимое test.txt
func readFrom() {
	file, err := os.Create("test2.txt")
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Panic(err)
		}
	}()

	origFile, err := os.Open("test.txt")
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := origFile.Close(); err != nil {
			log.Panic(err)
		}
	}()

	count, err := file.ReadFrom(origFile)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Copied", count, "bytes")
}

// ioCopy копирует содержимое из stdin в файл test.txt (но в текущей реализации логика некорректна)
func ioCopy() {
	file, err := os.OpenFile("test.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Здесь логика неверная - мы пытаемся копировать из stdout в файл,
	// но обычно нужно наоборот. Возможно, вы хотели копировать из файла в stdout?
	count, err := io.Copy(file, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Copied", count, "bytes from stdin to file")
}

// createFileBufio создает файл test3.txt с использованием буферизованной записи
func createFileBufio() {
	start := time.Now()
	file, err := os.Create("test3.txt")
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	defer func() {
		// Сначала сбрасываем буфер, потом закрываем файл
		if err := writer.Flush(); err != nil {
			log.Panic(err)
		}
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for i := 0; i < 100; i++ {
		if _, err = writer.WriteString(fmt.Sprintf("%d\n", i)); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("createFileBufio duration:", time.Since(start))
}
