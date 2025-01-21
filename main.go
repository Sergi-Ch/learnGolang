package main

import "fmt"

func main() {
	//variadicFunctions()
	//convertToArrayPointer()
	//passToFunction()
}

func variadicFunctions() {
	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 5, 6, 7)
	firstSlice := []int{9, 8, 7, 6, 5}
	secondSlice := []int{3, 3, 3, 3, 3, 3, 3}

	showAllElements(firstSlice...)
	newSlice := append(firstSlice, secondSlice...)
	fmt.Printf("type: %T value: %#v", newSlice, newSlice)
}

func showAllElements(values ...int) {
	for _, val := range values {
		fmt.Printf("Value: %#v\n", val)
	}
	fmt.Println()
}

func convertToArrayPointer() {
	initialSlice := []int{1, 2}
	fmt.Printf("type: %T value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("len: %d capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	intArray := (*[2]int)(initialSlice)
	fmt.Printf("type: %T value: %#v\n", intArray, intArray)
	fmt.Printf("len: %d capacity: %d\n\n", len(intArray), cap(intArray))

}

func passToFunction() {
	inicialSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("type: %T value: %#v\n", inicialSlice, inicialSlice)
	fmt.Printf("len: %d capacity: %d\n\n", len(inicialSlice), cap(inicialSlice))

	forTestPassSlice(inicialSlice)

	fmt.Printf("type: %T value: %#v\n", inicialSlice, inicialSlice)
	fmt.Printf("len: %d capacity: %d\n\n", len(inicialSlice), cap(inicialSlice))

	inicialSlice = append(inicialSlice, 3)
	fmt.Printf("type: %T value: %#v\n", inicialSlice, inicialSlice)
	fmt.Printf("len: %d capacity: %d\n\n", len(inicialSlice), cap(inicialSlice))

	inicialSlice = appendVaulue(inicialSlice)
	fmt.Printf("type: %T value: %#v\n", inicialSlice, inicialSlice)
	fmt.Printf("len: %d capacity: %d\n\n", len(inicialSlice), cap(inicialSlice))

}

func forTestPassSlice(slice []int) {
	slice[1] = 0
}

func appendVaulue(slice []int) []int {
	slice = append(slice, 10, 11)
	fmt.Printf("type: %T value: %#v\n", slice, slice)
	fmt.Printf("len: %d capacity: %d\n\n", len(slice), cap(slice))
	return slice
}
