package main
// a package is composed of files from the same directory and functions like a link to share resources

import (
	// "errors"
	"fmt" // from standard lib
	"github.com/jou-code/go-basics/assets" // import package by "module-name/package-name"
	// module name is usually repo address as a convention because of go tools that search for the link to download the code locally when developing
)

// functions don't need to be in order

func main() {

}

type person struct {
	name string
	age int32
	dogs []string
}

func person() {
	p1 := person {
		name: "drago",
		age: 24,
		dogs: nil,
	}
	p2 := person {
		name: "joana",
		age: 22,
		dogs: nil,
	}
}

func errorReturns() {
	

}

func maps() {
	// var hashmap map[string]string = map[string]string{}
	// need to initialize map ALWAYS, otherwise > segfault
	// basic hashmap:
	hashmap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	hashmap["key3"] = "value3"
	// hashmap with any type:
	hashmapAny := map[string]any{
		"key1": "value1",
		"key2": 2,
		"key3": map[string]string {
			"intern_key1": "intern_value1",
		},
	}
	hashmapAny["key4"] = "value4"
	// if hashmapAny["key2"] = "value4" > it overrides "value2"
	// loops + maps!
	for key, value := range hashmapAny {
		fmt.Println(key)
		fmt.Println(value)
	}

}

func loops() {
	arr := []string{"print", "these", "strings"}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	// range: keyword used on loops that iterates through collections and get values
	for index, str := range arr {
		fmt.Println(index)
		fmt.Println(str)
	}
	for index := range arr {
		fmt.Println(index)
	}
	for _, str := range arr {
		fmt.Println(str)
	}
}

func ifcond() {
	a := 10

	assets.Test()
	if b := 5; a > b {
		assets.Test()
		fmt.Println("a is bigger than b")
	} else {
		assets.Test()
		fmt.Println("a is smaller than b")
	}
}

func types() {
	// types: numerics(runes included), booleans and strings
	var name string
	type_inference := "value"

	var num byte
	var num uint8

	var num rune
	var num int32

	var ok bool
	bool := true
}

func arrays() {
	var runeArray [4]rune
	// ['a', 'b', 'c', 'd']

	var strArray [4]string
	// ["string1", "string2", "string3", "string4"]

	var intArray [4]*int8
	// [&var1, &var2, &var3, &var4]
}

func slice() {
	// slice is a dynamic array (reallocs new size when reaching full capacity)
	var x []int
	x := []int{} // {} to initialize value (empty, in this case)
	x := []int{1, 2, 3}

	// varName := value
	// := is a variable declaration and inicialization in a single line
	// go sets the var type by value setted

	//in slice cases it's necessary to set type
	// NOTE: search why

	a := []int{}
	a := append(a, 10, 20, 30)
}
*/
