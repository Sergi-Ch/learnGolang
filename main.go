package main

import "fmt"

func main() {
	var defaultMap map[int]string
	fmt.Printf("value: %#v tupe: %T\n", defaultMap, defaultMap)

	mapWithNew := *new(map[int64]string)
	fmt.Printf("value: %#v tupe: %T\n", mapWithNew, mapWithNew)

	mapByMake := make(map[string]string)
	fmt.Printf("value: %#v tupe: %T\n", mapByMake, mapByMake)

	mapByLiteral := map[int]string{1: "first", 2: "second"}
	fmt.Printf("value: %#v tupe: %T\n", mapByLiteral, mapByLiteral)

	fmt.Println(mapByLiteral[1])

	value, ok := mapByLiteral[3]
	fmt.Printf("value: %#v isThereInMap?>> %t\n", value, ok)

	mapByLiteral[3] = "Third"
	fmt.Printf("value: %#v tupe: %T\n", mapByLiteral, mapByLiteral)

	mapByLiteral[3] = "surprise"
	fmt.Printf("value: %#v tupe: %T\n", mapByLiteral, mapByLiteral)

	delete(mapByLiteral, 3)
	fmt.Printf("value: %#v tupe: %T\n", mapByLiteral, mapByLiteral)

	mapByLiteral[3] = "third"
	mapByLiteral[4] = "thou"

	for key, val := range mapByLiteral {
		fmt.Printf("key: %d value: %#v \n", key, val)
	}

}
