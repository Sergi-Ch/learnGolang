package main

import (
	"fmt"
	"math"
)

// Определение интерфейса Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Структура Circle
type Circle struct {
	Radius float64
}

// Реализация методов интерфейса Shape для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Структура Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Реализация методов интерфейса Shape для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Функция для вывода информации о фигуре
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

func interfaceValue() {
	fmt.Println("func interfaceValue begin")
	var testInterface Shape
	fmt.Printf("Type default testIntarface: %T value: %#v \n", testInterface, testInterface)
	if testInterface == nil {
		fmt.Printf("testInterface is nil \n")
	}
	var unTestInterface *Rectangle
	testInterface = unTestInterface
	fmt.Printf("Type testIntarface: %T value: %#v \n", testInterface, testInterface)
	if testInterface == nil {
		fmt.Printf("testInterface is nil \n")
	} else {
		fmt.Printf("it isn't nil \n")
	}
	namedTestIinteface := &Rectangle{Width: 5, Height: 10}
	fmt.Printf("Type namedTestIntarface: %T value: %#v \n", namedTestIinteface, namedTestIinteface)
	var perim float64 = namedTestIinteface.Perimeter()
	fmt.Printf("вызванный метод периметра у namedTestInterdace %f", perim)

}

func main() {
	// Создаем экземпляры Circle и Rectangle
	circle := Circle{Radius: 5}
	bigCircle := Circle{Radius: 10}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Выводим информацию о фигурах
	PrintShapeInfo(circle)
	PrintShapeInfo(bigCircle)
	PrintShapeInfo(rectangle)
	interfaceValue()
	// про утверждение типа
	var s Shape

	// Присваиваем значение типа Circle переменной s
	s = circle

	// Используем тип утверждение для доступа к конкретному типу
	if c, ok := s.(Circle); ok {
		fmt.Println("Circle Radius:", c.Radius)
	} else {
		fmt.Println("s is not a Circle")
	}

	// Присваиваем значение типа Rectangle переменной s
	s = rectangle

	// Используем тип утверждение для доступа к конкретному типу
	if r, ok := s.(Rectangle); ok {
		fmt.Println("Rectangle Width:", r.Width)
		fmt.Println("Rectangle Height:", r.Height)
	} else {
		fmt.Println("s is not a Rectangle")
	}

	//using switch
	s = circle

	// Используем type switch для обработки различных типов
	switch v := s.(type) {
	case Circle:
		fmt.Println("Circle Radius:", v.Radius)
	case Rectangle:
		fmt.Println("Rectangle Width:", v.Width)
		fmt.Println("Rectangle Height:", v.Height)
	default:
		fmt.Println("Unknown type")
	}

	// Присваиваем значение типа Rectangle переменной s
	s = rectangle

	// Используем type switch для обработки различных типов
	switch v := s.(type) {
	case Circle:
		fmt.Println("Circle Radius:", v.Radius)
	case Rectangle:
		fmt.Println("Rectangle Width:", v.Width)
		fmt.Println("Rectangle Height:", v.Height)
	default:
		fmt.Println("Unknown type")
	}

}
