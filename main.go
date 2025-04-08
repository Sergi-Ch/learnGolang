package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	//chanAsPromise()    // Демонстрация использования каналов как Promise
	//chanAsMutex()      // Демонстрация использования канала как Mutex
	//WithOutErrorGroup() // Реализация группы горутин без errgroup
	errGroup() // Реализация группы горутин с errgroup
}

// makeRequest имитирует асинхронный HTTP-запрос
// num - номер запроса для идентификации
// Возвращает канал, в который будет отправлен результат
func makeRequest(num int) <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second) // Имитация долгого запроса
		responseChan <- fmt.Sprintf("response numb %d", num)
	}()
	return responseChan
}

// chanAsPromise демонстрирует паттерн "как Promise" в Go
// Запускает два асинхронных запроса и ожидает их завершения
func chanAsPromise() {
	firstResponceChan := makeRequest(1)  // Первый асинхронный запрос
	SecondResponceChan := makeRequest(2) // Второй асинхронный запрос

	// Можно выполнять другие операции, пока запросы выполняются
	fmt.Println("non blocking")

	// Ожидаем завершения обоих запросов (аналог Promise.all)
	fmt.Println(<-firstResponceChan, <-SecondResponceChan)
}

// chanAsMutex демонстрирует использование канала как мьютекса
// для защиты общего ресурса (counter) от гонки данных
func chanAsMutex() {
	var counter int
	// Канал с буфером 1 используется как бинарный семафор
	mutexChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mutexChan <- struct{}{} // Захват "мьютекса"
			counter++               // Критическая секция
			<-mutexChan             // Освобождение "мьютекса"
		}()
	}

	wg.Wait()
	fmt.Println(counter) // Должно быть 1000
}

// WithOutErrorGroup показывает реализацию группы горутин
// с обработкой ошибок без использования errgroup
func WithOutErrorGroup() {
	var err error

	// Создаем отменяемый контекст для прерывания горутин
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(3) // Ожидаем завершения 3 горутин

	// Горутина 1
	go func() {
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done(): // Проверяем отмену
			return
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
		}
	}()

	// Горутина 2 (возвращает ошибку)
	go func() {
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("second started")
			err = fmt.Errorf("any error") // Устанавливаем ошибку
			cancel()                      // Отменяем другие горутины
		}
	}()

	// Горутина 3
	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done(): // Будет отменена из-за ошибки в горутине 2
			return
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println(err) // Выводим ошибку (если была)
}

// errGroup демонстрирует использование errgroup для управления группой горутин
// с автоматической отменой при возникновении ошибки
func errGroup() {
	// Создаем группу с общим контекстом
	g, ctx := errgroup.WithContext(context.Background())

	// Горутина 1
	g.Go(func() error {
		time.Sleep(time.Second)

		select {
		case <-ctx.Done(): // Проверяем отмену контекста
			return nil
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
			return nil
		}
	})

	// Горутина 2 (возвращает ошибку)
	g.Go(func() error {
		fmt.Println("second started")
		return fmt.Errorf("unexpeted error in request 2")
	})

	// Горутина 3
	g.Go(func() error {
		select {
		case <-ctx.Done(): // Будет отменена из-за ошибки в горутине 2
			return nil
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
			return nil
		}
	})

	// Ожидаем завершения всех горутин
	// Если была ошибка - выводим ее
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
