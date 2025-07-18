package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
)

type FileStats struct {
	ByteCount      int
	LineCount      int
	WordCount      int
	CharacterCount int
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

	whitespace_characters := []rune{0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x20}

	in_word := false

	for {
		ch, size, err := br.ReadRune()

		if err != nil && errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}

		fileInformation.CharacterCount++

		fileInformation.ByteCount += size

		if rune(ch) == rune(10) {
			fileInformation.LineCount++
		}

		if !in_word && !slices.Contains(whitespace_characters, ch) {
			in_word = true
			continue
		}

		if in_word && slices.Contains(whitespace_characters, ch) {
			in_word = false
			fileInformation.WordCount++
		}
	}

	return fileInformation
}

func EmitResult(code string, info FileStats, fileName string) {

	switch code {
	case "-c":
		fmt.Printf("%d %s\n", info.ByteCount, fileName)
	case "-l":
		fmt.Printf("%d %s\n", info.LineCount, fileName)
	case "-w":
		fmt.Printf("%d %s\n", info.WordCount, fileName)
	case "-m":
		fmt.Printf("%d %s\n", info.CharacterCount, fileName)
	}

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

	code := args[0]

	EmitResult(code, fileInformation, fileName)
}
