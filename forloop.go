package main

import "fmt"

func main() {
	n := 1
	fmt.Print("Masukkan N : ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
