package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := 3

	var data = [...]string{"Naufal", "Taufiq", "Ridwan"}
	fmt.Println(data)

	for i := 0; i < n; i++ {
		fmt.Print(i+1, ". Masukkan Nama : ")
		scanner.Scan()
		data[i] = scanner.Text()
	}
	fmt.Println("Menampilkan Data : ")
	for i := 0; i < n; i++ {
		fmt.Print(i+1, ". Nama : ", data[i]+"\n")
	}
}
