package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)
	fmt.Println("1.01 - 0.99 =", 1.01-0.99)

	fmt.Println(1 % 2)
	fmt.Println(2 % 2)
	fmt.Println(3 % 2)

	fmt.Println("length =", len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")

	character := "Hello, World"[1]
	fmt.Println("character=", character)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)


}
