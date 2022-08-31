package main

import (
	"fmt"
)

type Barang struct {
	name  string
	harga int
}

func main() {
	barang := [10]Barang{
		{"Benih Lele", 50000},
		{"Pakan lele cap menara", 25000},
		{"Probiotik A", 75000},
		{"Probiotik Nila B", 10000},
		{"Pakan Nila", 20000},
		{"Benih Nila", 20000},
		{"Cupang", 5000},
		{"Benih Nila", 30000},
		{"Benih Cupang", 10000},
		{"Probiotik B", 10000},
	}

	fmt.Println("Item yang memiliki harga sama: ")
	fmt.Println(HargaSama(barang[:], 10000))

	fmt.Println("Memanggil harga termurah dan termahal: ")
	fmt.Println(MurahMahal(barang[:]))

	fmt.Println("Total barang yang bisa dibeli: ")
	//fmt.Println(TotalP("Belum bisa"))
}

// func TotalP(harga int) ([]Barang, int) {

// }

func MurahMahal(arr []Barang) (Barang, Barang) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			var temp Barang
			if arr[i].harga > arr[j].harga {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr[0], arr[len(arr)-1]
}

func HargaSama(arr []Barang, harga int) []string {
	Panjang := len(arr)
	var h []string
	for i := 0; i < Panjang; i++ {
		if arr[i].harga == harga {
			h = append(h, arr[i].name)
		}
	}
	return h
}
