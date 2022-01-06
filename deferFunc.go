package main

import "fmt"

func selesai() {
	fmt.Println("Program selesai.")
}

func sayHello(name string) {
	defer selesai()
	fmt.Println("Hello", name)
}

func main() {
	sayHello("Naufal Taufiq Ridwan")
}
