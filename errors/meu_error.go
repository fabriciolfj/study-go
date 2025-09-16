package errors

import (
	"errors"
	"fmt"
)

var ErrorDivideByZero = errors.New("divide by zero")

func ExecuteMeuError() {
	fmt.Println("divide 1 by 0")
	_, err := precheckDivide(1, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("divide 2 by 0")
	divide(2, 0)
}

func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrorDivideByZero
	}

	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}
