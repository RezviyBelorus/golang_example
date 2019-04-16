package main

import "fmt"

const URL string = "http://golang.org"

var helloString string = "Hello world"

const (
	a = 5
	b = 10
	c = 15
)

func main() {

	fmt.Println(c)
	y := "shorter type"

	fmt.Println(helloString)
	fmt.Println("y =", y)

	var floatVar float32 = 32
	var intVar int32 = 32

	intVarShorten := 25600

	fmt.Println(floatVar)
	fmt.Println(intVar)
	fmt.Println(intVarShorten)

	fmt.Println("====From display func====")
	display(helloString)
}

func display(msg string) {
	fmt.Println(msg)
}
