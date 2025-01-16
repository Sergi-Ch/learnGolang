package main

import "fmt"

// Определение интерфейса Builder с методом Build
type Builder interface {
	Build()
}

// Тип Person, который будет встраиваться в другие типы
type Person struct {
	Name string
	Age  int
}

// Тип WoodBuilder, который встраивает тип Person
type WoodBuilder struct {
	Person // Встраивание типа Person
}

// Реализация метода Build для типа WoodBuilder
func (wb WoodBuilder) Build() {
	fmt.Println("Строю дом из дерева")
}

// Тип BrickBuilder, который встраивает тип Person
type BrickBuilder struct {
	Person // Встраивание типа Person
}

// Реализация метода Build для типа BrickBuilder
func (bb BrickBuilder) Build() {
	fmt.Println("Строю дом из кирпича")
}

// Тип Building, который встраивает интерфейс Builder
type Building struct {
	Builder // Встраивание интерфейса Builder
	Name    string
}

/*
// Пример метода для типа Person (закомментирован)
func (p Person) printName() {
    fmt.Println(p.Name)
}

// Пример функции для объяснения встраивания (закомментирован)
func explanation() {
    builder := WoodBuilder{Person{Name: "Вася", Age: 12}, "Боб"}
    fmt.Printf("Type: %T, Value: %#v \n", builder, builder)
    fmt.Println(builder.Person.Age)
    fmt.Println(builder.Age)

    builder.printName()
    fmt.Println(builder.Name)
}
*/

// Функция useCase демонстрирует использование встраивания и интерфейсов
func useCase() {
	// Создание экземпляра Building с WoodBuilder
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			Name: "Вася",
			Age:  40,
		}},
		Name: "Деревянная изба",
	}
	// Вызов метода Build для WoodBuilder
	woodenBuilding.Build()

	// Создание экземпляра Building с BrickBuilder
	brickBuilding := Building{
		Builder: BrickBuilder{
			Person{
				Name: "Петя",
				Age:  30,
			},
		},
		Name: "Кирпичный дом",
	}
	// Вызов метода Build для BrickBuilder
	brickBuilding.Build()
}

func main() {
	// Вызов функции useCase для демонстрации
	useCase()
}
