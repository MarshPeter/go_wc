package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type FileStats struct {
	ByteCount int
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		panic("You must include a file or flags to read from")
	}

	f, err := os.Open("./test_file.txt")

	if err != nil {
		panic(err)
	}

	br := bufio.NewReader(f)

	fileInformation := FileStats{}

	for {
		_, err := br.ReadByte()

		if err != nil && errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}

		fileInformation.ByteCount += 1

	}

	fmt.Printf("%d %s\n", fileInformation.ByteCount, "./test_file.txt")
}
