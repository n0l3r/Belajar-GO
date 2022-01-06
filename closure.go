package main

import "fmt"

func main() {
	count := 0
	sayHello := func() {
		fmt.Println("Halo Naufal")
		count++
	}

	sayHello()
	sayHello()
	fmt.Println("Total :", count)
}
