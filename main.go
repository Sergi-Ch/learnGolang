package main

import "fmt"

func main() {
	flag := 21
	if isOd := isOdd(flag); isOd == true {
		fmt.Println("it's odd")

	} else {
		fmt.Println("it isn't odd")
	}

}
func isOdd(flag int) bool {
	if flag%2 != 0 {
		return true
	} else {
		return false
	}
}
