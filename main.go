package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "it's loop for ")
	}
	count := 0
	for count < 10 {
		fmt.Println(count, "It's loop 'while' ")
		count++
	}
}
