package main

import (
	"errors"
	"fmt"
)

const anyPrice = 100

func main() {
	change, err := canYouBye(90, 17)
	checkError(change, err)
}

func canYouBye(money, age int) (int, error) {
	if anyPrice <= money && age >= 18 {
		return money - anyPrice, nil
	} else if age < 18 {
		return money, errors.New("you're child :( ")
	} else {
		return money, errors.New("it's expensive for you")
	}
}
func checkError(change int, err error) {
	if err != nil {
		fmt.Println("Неудача>>> ", err.Error())
	} else {
		fmt.Println("your change >>> ", change)
	}
}
