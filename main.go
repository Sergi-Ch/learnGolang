package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//withoutWait()
	//withWait()
	//wrongAdd()
	writeWithoutConcurrent()
	writeWithConcurrent()
	writeWithMutexConcurrent()
	readWithMutex()
	readWithRWMutex()
}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

func withWait() {
	var wg sync.WaitGroup //создаем wait группу
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i + 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("exit")
}

func wrongAdd() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {

		go func(i int) {
			wg.Add(1)

			defer wg.Done()
			fmt.Println(i + 1)
		}(i)

	}

}

func writeWithoutConcurrent() {
	fmt.Println("writeWithoutConcurrent")
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithConcurrent() {
	fmt.Println("\nwriteWithConcurrent")
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++

		}()

	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithMutexConcurrent() {
	fmt.Println("\nwriteWithMutexConcurrent")
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mu.Lock()
			counter++
			mu.Unlock()

		}()

	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	fmt.Println("\n readWithMutex")

	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			_ = counter //имитация чтения в переменную
			mu.Unlock()
		}()

	}

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithRWMutex() {
	fmt.Println("\n readWithRWMutex")

	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.RWMutex

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mu.RLock()
			_ = counter //имитация чтения в переменную
			mu.RUnlock()
		}()

	}

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
