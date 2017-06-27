package main

import (
	"fmt"
	"io"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func sampleReadFromString() {
	data, _ := ReadFrom(string.NewReader("from string"), 12)
	fmt.Println(data)
}

func main() {

}
