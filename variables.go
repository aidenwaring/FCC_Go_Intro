package main

import (
	"fmt"
	"strconv"
)

/*
	Levels of scope in Go:
	lowercase -> package level scope
	uppercase -> exported from package, globally visible
	block scope visibilty
*/

var a int
var i int = 42
var AgeOfPerson int = 42

// Block var declaration
var (
	name string = "John"
	age int = 21
	awesome bool = true
)

func output() {
	a = 2

	j := 99 // Can only used in block scope


	fmt.Println(a)
	fmt.Printf("%v, %T", i, i)
	fmt.Printf("%v, %T", j, j)

	fmt.Println(name, age, awesome)
	fmt.Println(convertingFloatToInt())
	fmt.Println(convertIntToString())
}

func convertingFloatToInt() int {
	var score float32 = 12.5

	// Explicit conversion required - points = score will not work
	var points = int(score)

	return points
}

func convertIntToString() string {
	var num int = 1

	// var name = string(num) -- Error, must use the string conversion package
	var word string = strconv.Itoa(num)
	return word
}