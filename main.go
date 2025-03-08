package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//baseSelect()
	//gracefulShutdown()
}

func baseSelect() {
	/*	bufferedChan := make(chan string, 3) //2
			bufferedChan <- "first"

			select {
			case str := <-bufferedChan:
				fmt.Println("read", str)
			case bufferedChan <- "second":
				fmt.Println("write", <-bufferedChan, <-bufferedChan)
			}

				unbufChan := make(chan int)

				go func() {
					time.Sleep(time.Second)
					unbufChan <- 1
				}()

					select {
					//case bufferedChan <- "third":
					//	fmt.Println("неблокирующая запись")
					case val := <-unbufChan:
						fmt.Println("блокирующее чтение", val)
					case <-time.After(time.Millisecond * 1500):
						fmt.Println("ожидание вышло")
						//default:
						//	fmt.Println("дефолтная ветка")
					}

		resultChan := make(chan int)
		timer := time.After(time.Second)

		go func() {
			defer close(resultChan)

			for i := 0; i < 1000; i++ {
				select {
				case <-timer:
					fmt.Println("время вышло")
					return
				default:
					time.Sleep(time.Millisecond * 10)
					resultChan <- i
				}
			}
		}()

		for v := range resultChan {
			fmt.Println(v)
		}*/
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	select {
	case <-timer:
		fmt.Println("Время вышло")
		return
	case sig := <-sigChan:
		fmt.Println("Остановилось сигналом: ", sig)
		return
	}

}
