package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	a, b := "Hi,", "there?"
	fmt.Println("Welcome to Go!")
	fmt.Println("The Square root of 81 is", math.Sqrt(81))
	fmt.Println("The random number between 1 - 99 is", rand.Intn(99))
	calc()
	fmt.Println(ConcatWithMultipleReturn(a, b))
	fmt.Println("The output of ConcatWithSingleReturn of "+a+" & "+b+" is ",
		ConcatWithSingleReturn(a, b))
	PointerDemo()
}

//PointerDemo - Pointer Demo Function
func PointerDemo() {
	a := 25
	b := &a
	fmt.Println("The Location of a is", b)
	fmt.Println("The value of a is", *b)
	fmt.Println("The value of a is", a)
	*b = 100
	fmt.Println("The value of a is", a)
	fmt.Println("The value of a is", *b)
}

//ConcatWithMultipleReturn - 2 values returned from function below
func ConcatWithMultipleReturn(a, b string) (string, string) {
	return a, b
}

//ConcatWithSingleReturn - 1 value returned from function below
func ConcatWithSingleReturn(a, b string) string {
	return a + " " + b
}

func calc() {
	a := 1.2
	b := float64(a)
	fmt.Println("The sum of two numbers is", add(a, b))
	fmt.Println("The difference of two numbers is", diff(a, b))
	fmt.Println("The product of two numbers is", product(a, b))
	fmt.Println("The quotient of two numbers is", divide(a, b))
}

func add(x, y float64) float64 {
	return x + y
}

func diff(x, y float64) float64 {
	return x - y
}

func product(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}
