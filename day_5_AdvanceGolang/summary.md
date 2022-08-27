# Summary Day 5

Pada hari ini kita sudah memasuki materi Golang advance, yang mana sudah mulai membahas hal-hal yang lebih kompleks, kita belajar beberapa materi, yang diantaranya yaitu:

## Pointer

Pointer ini digunakan sebagai penanda alamat memori yang ada pada variabel lain.
Sebagai contoh kita mempunyai variable A bertipe data int dengan nilai = 50, lalu kita mendeklarasikan variable pointer baru, yaitu variable B yang merujuk pada alamat variable A.
Misal pada saat A memiliki alamat memori 0x123 dengan nilai yang sudah ditentukan tadi yaitu 50, maka pada saat kita memanggil variable pointer B, maka yang terpanggil adalah alamat memori A tadi, dengan nilai 0x123.

## Method

Method ini gampangnya merupakan implementasi daripada struct, misal:

```
Type students struct{
    name string
}

func (s students) sayHello(){
    fmt.Println("Hello)
}
```

Nah pada contoh diatas, function sayHello merupakan method daripada struct student itu sendiri, sebanyak apapun implementasi turunan struct student, ia akan tetap disebut method.

## Public and Private

Dalam Go, penggunaan public dan private sangat mudah, cukup membedakan huruf awal daripada function yang akan kita deklarasikan, jika ia Public maka harus diawali huruf kapital, dan sebaliknya jika ia private maka diawali dengan huruf kecil.

## Interface

Interface ini gampangnya digunakan sebagai pengelompokan method-method.

## GoRoutine

Goroutine ini mirip sekali seperti thread, yang membuatnya spesial adalah, pada saat kita menggunakan ini, kita bisa menjalankan beberapa core processor sekaligus, sehingga program kita akan berjalan lebih cepat.
