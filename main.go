package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./test_file.txt")

	if err != nil {
		panic(err)
	}

	br := bufio.NewReader(f)

	for {
		b, err := br.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}

		fmt.Printf("%c", b)

		if err != nil {
			break
		}
	}

	fmt.Println("Hello, World")
}
