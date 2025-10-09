package simulacao

import (
	"errors"
	"fmt"
)

type Product struct {
	Id          int64
	Description string
	Price       float64
}

var InvalidPrice = errors.New("invalid price")

func Execute() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error handling product")
		}
	}()

	products := getProducts()

	for _, product := range products {
		fmt.Println(product)

		if product.Price > 10.0 {
			panic(InvalidPrice)
		}
	}
}

func getProducts() []Product {
	product := Product{
		Id:          1,
		Description: "Product 1",
		Price:       10.0,
	}

	product2 := Product{
		Id:          2,
		Description: "Product 2",
		Price:       10.0,
	}

	return []Product{product, product2}
}
