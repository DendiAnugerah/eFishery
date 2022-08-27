package main

import (
	"fmt"
	"sync"
)

type Student struct {
	name  string
	grade int
}

// ---------------- Method
func (s Student) SayHello() {
	fmt.Println("Hello", s.name)
}

// Merubah value dalam method
func (s *Student) ChangeName(name string) {
	s.name = name
}

// ---------------- Interface
type Repository interface {
	Hello() string
	GetDataPetaniById() string
}

type Petani struct {
	name string
	age  int
}

func (p Petani) Hello() string {
	return "Hello " + p.name
}

func main() {
	// // ---------------- Pointer
	// a := "Hello world"
	// fmt.Println("a: ", a)

	// b := &a
	// fmt.Println("b: ", b)
	// fmt.Println("b value: ", *b)

	// a = "Hello"
	// fmt.Println("new b value : ", *b)
	// fmt.Println("--------------------")
	// // ---------------- Method
	// student1 := Student{"Nurhuda", 13}
	// student1.SayHello()
	// student1.ChangeName("Dendu")
	// student1.SayHello()

	// // ---------------- Private and Public
	// fmt.Println("--------------------")
	// fmt.Println(library.SayHello("Dendi"))

	// // ---------------- Interface
	// fmt.Println("\n--------------------")
	// petani := Petani{
	// 	"Dendi",
	// 	12,
	// }
	// result := petani.Hello()
	// fmt.Println("Result: ", result)

	// ---------------- GoRoutines
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Hallo")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("world")
	}()

	func() {
		defer wg.Done()
		fmt.Println("dendi")
	}()
	wg.Wait()
}
