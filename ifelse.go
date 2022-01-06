package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string = "Naufal Taufiq Ridwan"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan Nama Langkap : ")
	scanner.Scan()

	if scanner.Text() == name {
		fmt.Println("Halo Naufal !")
	} else {
		fmt.Println("Kamu bukan Naufal")
	}

	// temp variabel di if else

	fmt.Print("Masukkan Nama Lengkap : ")
	scanner.Scan()

	if temp := scanner.Text(); temp == name {
		fmt.Println("Halo", temp)
	} else {
		fmt.Println("Kamu bukan", name+". Kamu adalah", temp)
	}
}
