// variable naming rules
// 1. letters, digits, -

package main

import "fmt" // this is already created package.

func variable() {
	/// variable declaration
	// var name string
	// var age int
	// var country string

	// name = "Mostafiz"
	// age = 22
	// country = "BD"

	var fullName = "Md Mostafiz" // it's understand fullName types. It's name is => Type Inferance
	name := "Mostafiz"           // sortly variable writing rule
	age := 22                    // sortly variable writing rule
	const country = "Bangladesh" // Not changeable

	fmt.Println(name)
	fmt.Printf("%v is %v years old", fullName, age)
	fmt.Println(fullName)
}
