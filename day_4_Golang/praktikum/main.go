package main

import (
	"fmt"
)

//struct
type Ikan struct {
	name    string
	spesies Kingdom
}

type Kingdom struct {
	famili string
	marga  string
	jenis  string
}

func main() {
	//------------ Struct ------------
	hewan := Ikan{
		name: "Nila",
		spesies: Kingdom{
			famili: "Chiclidae",
			marga:  "Oreochromis",
			jenis:  "Oreochromis niloticus",
		},
	}
	fmt.Println("---Struct---")
	fmt.Println("Nama Ikan: ", hewan.name)
	fmt.Println("Spesies  : ", hewan.spesies.famili)
	fmt.Println("Marga    : ", hewan.spesies.marga)
	fmt.Println("Jenis    : ", hewan.spesies.jenis)

	//Setelah diedit
	fmt.Println("\nSetelah diedit")
	hewan.spesies.jenis = "Ikan hilang"
	fmt.Println("Nama Ikan: ", hewan.name)
	fmt.Println("Spesies  : ", hewan.spesies.famili)
	fmt.Println("Marga    : ", hewan.spesies.marga)
	fmt.Println("Jenis    : ", hewan.spesies.jenis)

	//------------ Map ------------
	hewan1 := map[int]string{
		1: "Lumba-Lumba",
		2: "Hiu gila",
	}
	fmt.Println("\n---Map---")
	for key, value := range hewan1 {
		fmt.Println(key, ". Nama: ", value)
	}
	//Setelah diedit
	fmt.Println("\nSetelah diedit")
	hewan1[1] = "Ikan gila"
	for key, value := range hewan1 {
		fmt.Println(key, ". Nama: ", value)
	}

	//------------ Array ------------
	hewan2 := [...]Ikan{{
		name: "Nila",
		spesies: Kingdom{
			famili: "Chiclidae",
			marga:  "Oreochromis",
			jenis:  "Niloticus",
		},
	}, {name: "Hiu",
		spesies: Kingdom{
			famili: "Lamnidae",
			marga:  "Lamnifores",
			jenis:  "Hiu Stres",
		},
	}}
	fmt.Println("\n---Array---")
	fmt.Println(hewan2[0])
	fmt.Println(hewan2[1])
	//Setelah diedit
	hewan2[1] = Ikan{
		name: "nama aja sisanya kosong",
	}
	fmt.Println("\nSetelah diedit")
	fmt.Println(hewan2[0])
	fmt.Println(hewan2[1])

	//------------ Slice ------------
	fmt.Println("\n---Slice---")
	Hewan := []string{"Fish", "Cat", "Dog"}
	fmt.Println(Hewan)
	//Setelah diedit
	Hewan[1] = "Kucing"
	fmt.Println(Hewan)
}
