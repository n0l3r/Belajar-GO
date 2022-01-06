package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Print("Masukkan Nama : ")
	fmt.Scanf("%s", &name) // Space for submit

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan Nama Lengkap : ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Println(name)
}
