package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	myReader := io.Reader(strings.NewReader("Hello world"))
	arr := make([]byte, 5)
	for {
		n, err := myReader.Read(arr)
		//fmt.Printf("n = %d err = %v arr = %v\n", n, err, arr)
		fmt.Printf("%q\n", arr[:n])
		if err == io.EOF {
			break
		}
	}
}
