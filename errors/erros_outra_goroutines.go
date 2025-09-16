package errors

import (
	"errors"
	"time"

	"github.com/Masterminds/cookoo/safely"
)

func message() {
	println("Inside goroutine")
	panic(errors.New("Oops!"))
}

func Test() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(1000)
}
