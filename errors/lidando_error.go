package errors

import (
	"errors"
	"fmt"
)

func ExecuteTratamento() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("trapped panic : %s (%T)\n", err, err)
		}
	}()
	yikes()
}

func yikes() {
	panic(errors.New("oops"))
}
