package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(max-min) + min // рандомная переменная

	switch value {
	case 1:
		fmt.Println("Number is one")

	case 2, 3:
		fmt.Println("Number two or three")
		fallthrough // чтобы провалиться в следующий кейс

	case 10:
		fmt.Println("magic")

	default:
		fmt.Println("this is a default case, when other cases don't work")
	}

}
