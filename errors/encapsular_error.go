package errors

import (
	"errors"
	"fmt"
	"log"
)

func SendRequest2(req string) (string, error) {
	return "", fmt.Errorf("we got an error: %w", ErrTimeout)
}

func Execute3() {
	_, err := SendRequest2("Hello")
	if errors.Is(err, ErrTimeout) {
		log.Println("we got a timeout error")
	} else {
		log.Println("we got an error")
	}
}
