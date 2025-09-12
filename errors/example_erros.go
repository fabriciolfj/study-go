package errors

import (
	"errors"
	"fmt"
	"strings"
)

type error interface {
	Error() string
}

type ParseError struct {
	Message    string
	Line, Char int
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("Parse error at line %d, char %d: %s", p.Line, p.Char, p.Message)
}

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no parts")
	}

	return strings.Join(parts, ""), nil
}

func Execute() {
	result, err := Concat()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
