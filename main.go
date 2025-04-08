package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//Различные примеры работы с атомарными операциями и мьютексами
	//AddMutex()       // Пример с использованием мьютекса
	//AddAtomic()      // Пример с использованием atomic
	//StoreLoadSwap()  // Пример Load/Store/Swap операций
	//compareAndSwap() // Пример CompareAndSwap операции
	//atomicVal()      // Пример работы с atomic.Value
}

// AddMutex демонстрирует инкремент счетчика с использованием мьютекса
func AddMutex() {
	start := time.Now()
	var (
		counter int64          // Счетчик
		wg      sync.WaitGroup // Группа ожидания для горутин
		mu      sync.Mutex     // Мьютекс для синхронизации доступа
	)

	wg.Add(1000) // Добавляем 1000 задач в группу ожидания

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done() // Уменьшаем счетчик группы при завершении
			mu.Lock()       // Блокируем доступ к счетчику
			counter++       // Инкрементируем счетчик
			mu.Unlock()     // Разблокируем доступ
		}()
	}
	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println(counter)
	fmt.Println("with mutex: ", time.Now().Sub(start).Seconds())
}

// AddAtomic демонстрирует инкремент счетчика с использованием atomic
func AddAtomic() {
	start := time.Now()

	var (
		counter int64          // Счетчик
		wg      sync.WaitGroup // Группа ожидания
	)

	wg.Add(1000) // Добавляем 1000 задач

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done() // Уменьшаем счетчик группы
			// Атомарно увеличиваем счетчик
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait() // Ожидаем завершения
	fmt.Println(counter)
	fmt.Println("with atomic: ", time.Now().Sub(start).Seconds())
}

// StoreLoadSwap демонстрирует базовые атомарные операции
func StoreLoadSwap() {
	var counter int64

	// Атомарное чтение значения
	fmt.Println(atomic.LoadInt64(&counter))

	// Атомарная запись значения
	atomic.StoreInt64(&counter, 5)
	fmt.Println(atomic.LoadInt64(&counter))

	// Атомарная замена значения и возврат старого
	fmt.Println(atomic.SwapInt64(&counter, 10))
	// Проверка нового значения
	fmt.Println(atomic.LoadInt64(&counter))
}

// compareAndSwap демонстрирует операцию сравнения и замены
func compareAndSwap() {
	var (
		counter int64          // Счетчик
		wg      sync.WaitGroup // Группа ожидания
	)
	wg.Add(100) // 100 горутин

	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()

			// Пытаемся изменить значение, только если оно равно 0
			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return
			}
			// Выводим номер горутины, которой удалось изменить значение
			fmt.Println("swapped goroutine number is", i)
		}(i)
	}
	wg.Wait()
	fmt.Println(counter)
}

// atomicVal демонстрирует работу с atomic.Value
func atomicVal() {
	var (
		value atomic.Value // Значение произвольного типа
	)

	// Сохраняем значение
	value.Store(1)
	// Читаем значение
	fmt.Println(value.Load())

	// Пытаемся заменить значение, если текущее равно 1
	fmt.Println(value.CompareAndSwap(1, 3))
	// Читаем новое значение
	fmt.Println(value.Load())
}
