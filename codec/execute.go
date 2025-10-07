package codec

import (
	"fmt"

	"github.com/ugorji/go/codec"
)

func Execute() {
	jh := new(codec.JsonHandle)
	u := &User{
		Name:  "Inigo Montoya",
		Email: "inigo@montoya.example.com",
	}
	var out []byte
	err := codec.NewEncoderBytes(&out, jh).Encode(&u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
