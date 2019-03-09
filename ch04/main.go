package main

import "fmt"

var globalVar string = "Global"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = "hello "
	var y = "world"
	fmt.Println(x == y)

	x = "hello"
	y = "hello"
	fmt.Println(x == y)

	dogName := "Bingo"
	fmt.Println("My dog's name is ", dogName)

	fmt.Println(globalVar)

	const constante string = "Const"
	fmt.Println(constante)

	const (
		c1 = "var1"
		c2 = "var2"
		c3 = "var3"
	)

	fmt.Println(c1, c2, c3)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}

func f() {
	fmt.Println(globalVar)
}
