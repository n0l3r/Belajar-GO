package main

import "fmt"

func main() {
	var score float64
	var hurufMutu string

	fmt.Print("Masukkan Nilai : ")
	fmt.Scanf("%f", &score)

	switch {
	case score >= 80:
		hurufMutu = "A"
		fmt.Println("Nilai Kamu A")
	case score >= 75:
		hurufMutu = "AB"
		fmt.Println("Nilai Kamu AB")
	case score >= 60:
		hurufMutu = "B"
		fmt.Println("Nilai Kamu B")
	case score >= 50:
		hurufMutu = "C"
		fmt.Println("Nilai Kamu C")
	case score < 50:
		hurufMutu = "E"
		fmt.Println("Nilai Kamu E")
	}

	switch hurufMutu {
	case "E":
		fmt.Println("Kamu Tidak Lulus!")
	default:
		fmt.Println("Kamu Lulus!")
	}
}
