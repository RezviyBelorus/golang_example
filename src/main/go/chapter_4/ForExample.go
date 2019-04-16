package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("New loop")
	for i := 99; i <= 110; i++ {
		if i%2 == 0 {
			fmt.Println("even =", i)
		} else { // odd
			fmt.Println("odd =", i)
		}
	}

	fmt.Println("End if the loop")
}
