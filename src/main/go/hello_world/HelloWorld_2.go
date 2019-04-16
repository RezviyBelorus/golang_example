package main

import (
	"fmt"
	"os"
)

func main() {
	inputArgs := os.Args

	//in args[0] system var
	fmt.Println("Hello from there " + inputArgs[0])

	fmt.Println("Hello from there " + inputArgs[1])

	fmt.Println(0111)

}
