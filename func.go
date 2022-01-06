package main

import (
	"fmt"
	"os"
)

func printMenu() {
	fmt.Println("\n[+] Operasi [+]")
	fmt.Print("1. Penjumlahan\n2. Pengurangan\n3. Perkalian\n4. Pembagian\n\n")
}

func penjumlahan(a float64, b float64) float64 {
	return a + b
}

func pengurangan(a float64, b float64) float64 {
	return a - b
}

func perkalian(a float64, b float64) float64 {
	return a * b
}

func pembagian(a float64, b float64) float64 {
	return a / b
}

func main() {
	var operasi int8
	var num1, num2, result float64

	fmt.Print("Angka ke-1 : ")
	fmt.Scanf("%f", &num1)

	fmt.Print("Angka ke-2 : ")
	fmt.Scanf("%f", &num2)

	printMenu()
	fmt.Print("Pilih Menu : ")
	fmt.Scanf("%d", &operasi)

	switch operasi {
	case 1:
		result = penjumlahan(num1, num2)
	case 2:
		result = pengurangan(num1, num2)
	case 3:
		result = perkalian(num1, num2)
	case 4:
		result = pembagian(num1, num2)
	default:
		fmt.Println("Operasi Tidak Tersedia")
		os.Exit(0)
	}
	fmt.Println("Hasil =", result)
}
