package main

import "fmt"

func ganjil(start int, end int) {
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
}

func genap(start int, end int) {
	for i := start; i <= end; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i, " ")
	}
}

func main() {
	var menu int8
	var start, end int

	fmt.Println("[+] MENU [+]")
	fmt.Print("1. Ganjil\n2. Genap\n\n")

	fmt.Print("Pilih Menu : ")
	fmt.Scanf("%d", &menu)

	fmt.Print("Range a - b : ")
	fmt.Scanf("%d %d", &start, &end)

	switch menu {
	case 1:
		ganjil(start, end)
	case 2:
		genap(start, end)
	default:
		fmt.Println("Menu yang dimasukkan tidak sesuai!")
	}
}
