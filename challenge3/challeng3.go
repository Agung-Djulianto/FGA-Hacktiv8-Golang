package main

import (
	"fmt"
)

func main() {
	kalimat := "selamat datang bulan ramadhan 2023"

	// Hitung kemunculan setiap huruf
	mapHuruf := make(map[string]int)
	for _, h := range kalimat {
		mapHuruf[string(h)]++
	}

	// Tampilkan setiap huruf
	for _, h := range kalimat {
		fmt.Println(string(h))
	}

	// Tampilkan jumlah kemunculan setiap huruf
	fmt.Println(mapHuruf)
}
