package main

import "fmt"

type person struct {
	Age  int
	Name string
}

func explainArrays() {
	var intArr [5]int
	stringArr := [...]string{"first", "second", "third"}
	people := [...]person{
		{
			Age:  30,
			Name: "NameOne",
		},
		{
			Age:  25,
			Name: "NameTwo",
		},
	}

	for i := 0; i < len(intArr); i++ {
		fmt.Printf("index %d value: %d \n", i, intArr[i])
	}
	for i := 0; i < len(stringArr); i++ {
		fmt.Printf("index %d value: %s \n", i, stringArr[i])
	}
	for i := 0; i < len(people); i++ {
		fmt.Printf("index %d value: %#v \n", i, people[i])
	}

	fmt.Println("for-range")

	for ind, value := range people {
		fmt.Printf("index %d value: %#v \n", ind, value)
	}
	for _, value := range people {
		fmt.Printf("value: %#v \n", value)
	}
	for ind, _ := range people {
		fmt.Printf("index %d  \n", ind)
	}

	NewIntArr := changeArr(intArr)
	fmt.Println("after copping")
	fmt.Printf("Type %T value %#v \n", intArr, intArr)
	fmt.Printf("Type %T value %#v \n", NewIntArr, NewIntArr)

}
func changeArr(arr [5]int) [5]int {
	arr[4] = 8
	return arr
}

func explainSlices() {
	fmt.Println("about slices>>")
	var defaultSlice []int
	fmt.Printf("Type %T value %#v \n", defaultSlice, defaultSlice)
	fmt.Printf("length %d capacity %#v\n\n", len(defaultSlice), cap(defaultSlice))
	stringSliceLiteral := []string{"first", "second"}
	fmt.Printf("Type %T value %#v \n", stringSliceLiteral, stringSliceLiteral)
	fmt.Printf("length %d capacity %#v\n\n", len(stringSliceLiteral), cap(stringSliceLiteral))

	sliceByMake := make([]int, 3, 5)
	fmt.Printf("Type %T value %#v \n", sliceByMake, sliceByMake)
	fmt.Printf("length %d capacity %#v\n\n", len(sliceByMake), cap(sliceByMake))
	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5, 6)

	fmt.Printf("Type %T value %#v \n", sliceByMake, sliceByMake)
	fmt.Printf("length %d capacity %#v\n\n", len(sliceByMake), cap(sliceByMake))

	for ind, value := range sliceByMake {
		fmt.Printf("index: %d value: %#v\n", ind, value)
	}
}

func main() {
	explainArrays()

	explainSlices()
}
