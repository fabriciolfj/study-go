package metodo

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string `help:"name of the animal"`
}

type Home struct {
	string
}

func (h *Home) AddValue(value string) {
	h.string = value
}

func ChangeHome() {
	h := Home{
		"Hello",
	}
	fmt.Println(h.string)

	h.AddValue("World")
	fmt.Println(h.string)
}

func (a Animal) speak() string {
	switch a.name {
	case "cat":
		return "meow"
	case "dog":
		return "woof"
	default:
		if member, ok := reflect.TypeOf(a).FieldByName("name"); ok {
			return fmt.Sprintf("unknown %s", member.Tag.Get("help"))
		}
		return "unknown"
	}
}

func CreateNewAnimal() {
	a := Animal{name: "cat"}

	fmt.Println(a.speak())

	a.name = "dog"
	fmt.Println(a.speak())

	a.name = "bird"
	fmt.Println(a.speak())
}
