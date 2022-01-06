package main

import "fmt"

func getSayHello(name string) string {
	return "Halo, " + name
}
func main() {
	getSay := getSayHello
	result := getSay("Naufal")
	fmt.Println(result)
}
