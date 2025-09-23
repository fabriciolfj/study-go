package logs

import (
	"fmt"
	"runtime"
)

func Execute5() {
	foo()
}
func foo() {
	bar()
}
func bar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Trace:\n %s\n", buf)
}
