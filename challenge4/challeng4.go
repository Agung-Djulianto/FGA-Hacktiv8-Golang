package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	teman1 := Teman{"arik", "gresik kota baru", "programer", "ingin menambah keahlian"}
	teman2 := Teman{"wiwit", "porong kesamben", "mahasiswa", "Ingin meningkatkan skill programming"}
	teman3 := Teman{"izzi", "porong kesambi", "seles", "ingin menjadi backend golang"}
	teman4 := Teman{"andre", "sidoarjo", "mahasiswa", "tertarik dengan bahasa pemrograman golang"}
	teman5 := Teman{"candra", "tulangan sidoarjo", "mahasiswa", "meningkatkan skill pemrograman golang"}

	args := os.Args[1:]
	for _, arg := range args {
		switch arg {
		case "1":
			display(teman1)
		case "2":
			display(teman2)
		case "3":
			display(teman3)
		case "4":
			display(teman4)
		case "5":
			display(teman5)
		default:
			fmt.Println("Tidak ada data teman dengan nomor absen", arg)
		}
	}
}

func display(isi Teman) {
	fmt.Println("Nama:", isi.Nama)
	fmt.Println("Alamat:", isi.Alamat)
	fmt.Println("Pekerjaan:", isi.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", isi.Alasan)
	fmt.Println()
}
