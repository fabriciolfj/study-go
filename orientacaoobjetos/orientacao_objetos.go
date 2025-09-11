package orientacaoobjetos

import (
	"fmt"
)

type Animal interface {
	speak()
}

type Cat struct{}

func (c Cat) speak() {
	fmt.Println("meow")
}

func NewCat() *Cat {
	return &Cat{}
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("woof")
}

func NewDog() *Dog {
	return &Dog{}
}

type LLama struct{}

func NewLlama() *LLama {
	return &LLama{}
}

func Execute() {
	var a Animal

	c := NewCat()
	a = c
	a.speak()

	d := NewDog()
	a = d
	a.speak()

	//l := NewLlama()
	//a = l
}
