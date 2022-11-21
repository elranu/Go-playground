package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("Pepe")
	var myInt, _ = strconv.Atoi("123")
	str := strconv.Itoa(123)
	myInt += 1
	fmt.Println(myInt)
	fmt.Println(str)
	myBoolean, _ := strconv.ParseBool("1")
	fmt.Println(reflect.TypeOf(myBoolean))

	//array
	var exampleArray [3]int
	exampleArray2 := [3]int{1, 2, 3}
	fmt.Println("-------Array---------")
	fmt.Println(exampleArray[0])
	fmt.Println(exampleArray2[2])

	exampleSlice := []string{"hello", "world", "coco", "pepe", "cube"}
	fmt.Println("-------Slice---------")
	fmt.Println(exampleSlice[1])
	fmt.Printf("exampleSlice: %v \n", exampleSlice)

	firstTwo := exampleSlice[:2]
	lastThree := exampleSlice[2:4]
	lastElement := exampleSlice[len(exampleSlice)-1:]

	fmt.Printf("firstTwo %+v \n", firstTwo)
	fmt.Printf("lastThree %+v - Last Element: %v \n", lastThree, lastElement)

}
