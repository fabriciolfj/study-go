package reflection

import (
	"bytes"
	"fmt"
)

type Stringer interface {
	String() string
}

func Execute2() {
	b := bytes.NewBuffer([]byte("hello"))
	if isStringer(b) {
		fmt.Println("%T is a stringer\n", b)
	}

	i := 123
	if isStringer(i) {
		fmt.Println("%T is a stringer\n", i)
	}
}

func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}
