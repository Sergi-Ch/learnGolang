package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//baseKnowLedge()
	workerPool()
}

func baseKnowLedge() {
	ctx := context.Background()
	fmt.Println(ctx)
	ctxD := context.TODO()
	fmt.Println(ctxD)

	WithValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println(WithValue.Value("name"))

	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err())
	cancel()
	fmt.Println(withCancel.Err())

	withDeadLine, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	fmt.Println(withDeadLine.Deadline())
	fmt.Println(withDeadLine.Err())
	fmt.Println(<-withDeadLine.Done())

	withTimeOut, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	fmt.Println(<-withTimeOut.Done())
}

func workerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(ctx, time.Millisecond*3)

	defer cancel()

	wg := &sync.WaitGroup{}

	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			numbersToProcess <- i
		}

		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}

	fmt.Println(counter)

}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {

		case <-ctx.Done(): // выход если контекст завершился
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}

			time.Sleep(time.Millisecond) //типо долговременная обработка
			processed <- value * value

		}

	}
}
