package main

import (
	"fmt"
)

// Numbers - объединение типов для использования в дженериках
type Numbers interface {
	int64 | float64
}

// Numbrs - обобщенный тип срезов чисел (с опечаткой в названии)
type Numbrs[T Numbers] []T

func main() {
	//showSum()
	//showContains()
	//showAny()
	//unionInterfaceAndType()
}

// showSum демонстрирует функцию sum для разных типов чисел
func showSum() {
	floats := []float64{1.0, 2.0, 3.0}
	ints := []int64{1, 2, 3}
	fmt.Println(sum(floats))
	fmt.Println(sum(ints))
}

// sum возвращает сумму элементов среза (работает с int64 и float64)
func sum[V int64 | float64](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// showContains демонстрирует поиск элемента в срезах разных типов
func showContains() {
	type Person struct {
		name     string
		age      int64
		jobTitle string
	}

	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("int: ", contains(ints, 4))

	strings := []string{"Vasya", "Anton", "Katya"}
	fmt.Println("string:", contains(strings, "Katya"))
	fmt.Println("string:", contains(strings, "Sasha"))

	people := []Person{
		{
			name:     "Vasya",
			age:      20,
			jobTitle: "programmer",
		},
		{
			name:     "Dasha",
			age:      28,
			jobTitle: "Designer",
		},
	}

	// Демонстрация сравнения структур
	fmt.Println("structs: ", contains(people, Person{
		name:     "Vasya",
		age:      21,
		jobTitle: "programmer",
	}))

	fmt.Println("structs: ", contains(people, Person{
		name:     "Vasya",
		age:      20,
		jobTitle: "programmer",
	}))
}

// contains проверяет наличие элемента в срезе (для comparable типов)
func contains[T comparable](elements []T, searchEl T) bool {
	for _, el := range elements {
		if searchEl == el {
			return true
		}
	}
	return false
}

// show демонстрирует работу с любыми типами (any)
func showAny() {
	show(1, 2, 3)
	show("one", "two", "three")
	show([]int64{1, 2, 3}, []int64{4, 5, 6})
	show(map[string]int64{
		"first":  1,
		"second": 2,
	})
	show(interface{}(1), interface{}("string"), any(struct {
		name string
	}{name: "Vova"}))
}

// show выводит элементы любого типа
func show[T any](entities ...T) {
	fmt.Println(entities)
}

// unionInterfaceAndType демонстрирует использование интерфейса типов с дженериками
func unionInterfaceAndType() {
	var ints Numbrs[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...) // ... распаковывает срез

	floats := Numbrs[float64]{1.0, 2, 3, 4, 5}

	fmt.Println("ints:", ints)
	fmt.Println("floats:", floats)

	fmt.Println("sum ints:", sumUnionInterface(ints))
	fmt.Println("sum floats:", sumUnionInterface(floats))
}

// sumUnionInterface возвращает сумму для типов, определенных в Numbers
func sumUnionInterface[V Numbers](numbers []V) V {
	var sum V
	for _, numb := range numbers {
		sum += numb
	}
	return sum
}
