# Summary Day 4

Hari ini kita belajar mengenai basic dari bahasa pemrograman Go.

## Tipe Data

Golang sendiri memiliki banyak tipe data, dan standar yang digunakan adalah:

- Numerik Non-Desimal
- Numerik Desimal
- Boolean
- String

## Variable

Dalam pendeklarasian variabel, golang memiliki berbagai cara, diantaranya yaitu yang dituliskan tipe datanya, dan juga tidak.
| Contoh | Penggunaan |
| - | - |
| Variable beserta tipe data | `var <nama-variable> <tipe-data> = <nilai>` |
| Variable tanpa tipe data | `<nama-variable> := <nilai>`|
| Multi variable | `var <nama-1>, <nama-2>, <dst..> <tipe-data> = <nilai1>, <nilai2>, <dst>` |
| Variable underscore | `_ = <Sampah>` |
| Constant | `const <nama-variable> <tipe-data> = <nilai>` |

## Condition

Condition dalam golang juga sama halnya bahasa pemrograman yang lain.\
Contoh:

```
var point = 8

if point == 10 {
    fmt.Println("lulus dengan nilai sempurna")
} else if point > 5 {
    fmt.Println("lulus")
} else if point == 4 {
    fmt.Println("hampir lulus")
} else {
    fmt.Printf("tidak lulus. nilai anda %d\n", point)
}
```

## Looping

Perulangan digunakan untuk mengulang proses eksekusi blok tanpa henti, selama kondisi yang dijadikan acuan terpenuhi. Ada beberapa cara menggunakan for, antara lain:

```
// For range
fruits := [5]string {"Bayam", "Nanas"}
for i, fruit := range fruits{
    fmt.Println(fruit)
}

// For loop
for i := 0; i < 5; i++{
    fmt.Println("Angka", i)
}

// For loop break continue
for i := 0; i < 5; i++{
    if i == 1{
        continue
    } else if i = 3{
        break
    }
    fmt.Println("Angka", i)
}
```

## Function

Sebenernya kita sudah menuliskan fungsi pada saat awal menjalankan kode pertama kita, yaitu Fungsi main pada hello world, juga fungsi ini memiliki banyak jenisnya, salah satu contohnya yaitu:

```
// Multiple return value
func swap(x, y string) (string, string){
    return y, x
}

func main(){
    a, b := swap("Hello", "world")
    fmt.Println(a, b)

}
```

## Struct

Struct ini merupakan kumpulan variable atau fungsi yang dibungkus dalam tipe data baru dengan nama tertentu.\
Contoh deklarasi struct:

```
type student struct{
    name string
    grade int
}
```

Lalu kita juga dapat memasukkan struct dalam struct tersebut juga, sebagai contoh penerapannya.

```
type student struct{
    name string
    grade int
    school sekolah <- sekolah dipanggil pada struct di bawah
}

type sekolah struct{
    alamat string
    kelas int
}
```

## Data Struktur

### Map

Map ini memiliki key-value, dan setiap key harus bersifat unik, karena digunakan sebagai pengaksesan nantinya.\
Contoh penerapan:

```
chicken := map[string]int{
    "januari" = 1
    "februari" = 2
}

```

### Array

Array merupakan sekumpulan data beripe sama, cara pendeklarasian array pada golang cukup mudah.

```
var name [5]string
name[0] = "dendi"
```

## Slice

Slice ini merupakan reference elemen array, nah yang membedakan adalah jika jumlah elemen tidak dituliskan maka variable tersebut adalah slice.

```
var fruits = []string {"dendi", "oke"}
```

## Standar Library

Golang memiliki sangat banyak standar library yang kita bisa gunakan untuk mempermudah proses pengerjaan kita.

## Defer

Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.
