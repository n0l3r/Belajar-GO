package main

import "fmt"

func getFactorial(num int) int {
	if num < 1 {
		return 1
	}
	return num * getFactorial(num-1)
}

func main() {
	var num, result int
	fmt.Println("Faktorial menggunakan rekursif")
	fmt.Print("Masukkan N : ")
	fmt.Scanf("%d", &num)

	result = getFactorial(num)
	fmt.Println("Hasil :", result)
}
