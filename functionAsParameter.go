package main

import "fmt"

type Filter func(string) bool

func getSayHello(name string, filter Filter) string {
	var result string

	if filter(name) {
		result = "Halo, " + name
	} else {
		result = "Maaf kamu tidak dikenali."
	}
	return result
}

func cekNama(name string) bool {
	return name == "Naufal"
}

func main() {
	var name string
	fmt.Print("Masukkan Nama : ")
	fmt.Scanf("%s", &name)

	result := getSayHello(name, cekNama)
	fmt.Println(result)
}
