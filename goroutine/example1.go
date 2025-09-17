package goroutine

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Execute() {
	fmt.Println("Type anything below for up to 3 seconds")
	go echo(os.Stdin, os.Stdout)
	time.Sleep(3 * time.Second)
	fmt.Println("Time out")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
