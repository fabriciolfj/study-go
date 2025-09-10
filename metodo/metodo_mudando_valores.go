package metodo

import (
	"fmt"
)

type character struct {
	name string
}

func (c *character) fixName() {
	c.name = "John"
}

func AlterName() {
	c := character{name: "Inigo"}
	fmt.Println(c.name)
	c.fixName()
	fmt.Println(c.name)
}
