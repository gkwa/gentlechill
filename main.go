package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	myVar1 := "Hello"
	myVar2 := 42
	myVar3 := true

	str := spew.Sdump(myVar1, myVar2, myVar3)
	fmt.Println(str)
}

func test2() {
	nestedMap := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": map[string]interface{}{
			"nestedKey1": true,
			"nestedKey2": "value2",
		},
	}

	str := spew.Sdump(nestedMap)
	fmt.Println(str)
}

func test3() {
	type InnerStruct struct {
		NestedValue1 int
		NestedValue2 string
	}

	type OuterStruct struct {
		Value1 int
		Value2 string
		Nested InnerStruct
	}

	nestedStruct := InnerStruct{
		NestedValue1: 42,
		NestedValue2: "Hello",
	}

	outerStruct := OuterStruct{
		Value1: 10,
		Value2: "World",
		Nested: nestedStruct,
	}

	str := spew.Sdump(outerStruct)
	fmt.Println(str)
}
