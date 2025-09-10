package main

import "fmt"

func anonimo() {
	animal := struct {
		name  string
		speak func() string
	}{
		name: "dog",
		speak: func() string {
			return "woof"
		},
	}

	fmt.Println(fmt.Sprintf("%s says %s", animal.name, animal.speak()))
}
