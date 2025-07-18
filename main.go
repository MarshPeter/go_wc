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

func GetFile(fileName string) *os.File {
	f, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	return f
}

func GetContentStatistics(br *bufio.Reader) FileStats {
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

	return fileInformation
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		panic("You must include a file or flags to read from")
	}

	fileName := args[1]

	f := GetFile(fileName)

	br := bufio.NewReader(f)

	fileInformation := GetContentStatistics(br)

	fmt.Printf("%d %s\n", fileInformation.ByteCount, fileName)
}
