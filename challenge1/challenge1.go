package main

import "fmt"

func main() {

	var i = 21
	var kondisi bool = true
	var unik = 1103
	var kosong = ""
	var angka = 15
	var rusia = 'Я'
	var k float64 = 123.456
	var nilai8 uint = 21
	var nilai10 uint = 031

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% %s \n", kosong)
	fmt.Printf("%t \n", kondisi)
	fmt.Printf("%c \n", unik)
	fmt.Printf("%b \n", i)
	fmt.Println(nilai8)
	fmt.Printf("%d \n", nilai10)
	fmt.Printf("%x \n", angka)
	fmt.Printf("%X \n", angka)
	fmt.Printf("%U \n", rusia)
	fmt.Printf("%.6f \n%e", k, k)

}
