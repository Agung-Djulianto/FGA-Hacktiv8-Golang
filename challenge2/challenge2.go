package main

import "fmt"

func main() {

	var angka int
	var unik string = "CAШAPBO"

	for angka = 0; angka <= 10; angka++ {
		if angka <= 4 {
			fmt.Printf("Nilai i = %d\n", angka)
		}
	}

	for angka = 0; angka <= 10; angka++ {
		if angka == 5 {
			for i, c := range unik {
				fmt.Printf("character %#U start at byte position %d\n", c, i)
			}
		} else {
			fmt.Printf("Nilai j = %d\n", angka)
		}
	}
}
