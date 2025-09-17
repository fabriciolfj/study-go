package goroutine

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func Executar2() {
	var wg sync.WaitGroup
	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()
	fmt.Println("compressed %d files\n", len(os.Args[1:]))
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()
	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()
	return err
}
