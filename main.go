package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
)

const (
	BYTE_FLAG      string = "-c"
	LINE_FLAG      string = "-l"
	WORD_FLAG      string = "-w"
	CHARACTER_FLAG string = "-m"
	ALL_FLAG       string = "-c-l-w"
)

type TextStats struct {
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

func GetContentStatistics(br *bufio.Reader) TextStats {
	fileInformation := TextStats{}

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

func EmitResult(code string, info TextStats, fileName string) {

	switch code {
	case BYTE_FLAG:
		fmt.Printf("%d %s\n", info.ByteCount, fileName)
	case LINE_FLAG:
		fmt.Printf("%d %s\n", info.LineCount, fileName)
	case WORD_FLAG:
		fmt.Printf("%d %s\n", info.WordCount, fileName)
	case CHARACTER_FLAG:
		fmt.Printf("%d %s\n", info.CharacterCount, fileName)
	case ALL_FLAG:
		fmt.Printf("  %d %d %d %s\n", info.LineCount, info.WordCount, info.ByteCount, fileName)
	}
}

func IsFlag(str string) bool {
	return str == BYTE_FLAG ||
		str == LINE_FLAG ||
		str == WORD_FLAG ||
		str == CHARACTER_FLAG ||
		str == ALL_FLAG
}

func main() {
	args := os.Args[1:]

	var fileName string
	var code string
	isFileBased := true

	if len(args) == 0 {
		code = ALL_FLAG
		isFileBased = false
	}

	if len(args) == 1 {
		if IsFlag(args[0]) {
			code = args[0]
			isFileBased = false
		} else {
			fileName = args[0]
		}
	}

	if len(args) == 2 {
		code = args[0]
		fileName = args[1]
	}

	var br *bufio.Reader

	if isFileBased {
		f := GetFile(fileName)
		br = bufio.NewReader(f)
	} else {
		br = bufio.NewReader(os.Stdin)
	}

	textInformation := GetContentStatistics(br)

	if len(args) == 2 {
		EmitResult(code, textInformation, fileName)
	}

	if len(args) == 1 && isFileBased {
		EmitResult("-c-l-w", textInformation, fileName)
	}

	if len(args) < 2 && !isFileBased {
		EmitResult(code, textInformation, "")
	}
}
