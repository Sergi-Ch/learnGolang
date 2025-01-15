package main

import "fmt"

func main() {
	definition()
}

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("Периметр квадрата: %d \n", s.Side*4)

}

func (s *Square) Scale(multiplier int) { // умножение стороны квадрата на множитель
	fmt.Printf("%T %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T %#v \n", s, s)
}

func definition() {
	testStruct := Square{Side: 5}
	testPStruct := &testStruct
	secondTestStruct := Square{Side: 2}

	testStruct.Perimeter()
	secondTestStruct.Perimeter()

	testPStruct.Scale(7)
}
