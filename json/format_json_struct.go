package json

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name           string  `json:"animal_name"`
	ScientificName string  `json:"scientific_name"`
	Weight         float64 `json:"animal_average_weight"`
}

func Execute() {
	a := Animal{
		Name:           "cat",
		ScientificName: "felis catus",
		Weight:         10.0}

	output, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}
