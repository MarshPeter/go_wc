# go_wc

## Motivation

This is a program to explore the language of go, by roughly replicating the unix command wc. 

## Functionality

The following flags correctly work and match outputs of the wc command. 

"-c" - A count of all bytes in the text.

"-l" - A count of all lines in the text.

"-w" - A count of all words separated by whitespace that are in the text. 

"-m" - A count of all utf-8 characters in the text. 

## Usage

### Building the program

The following should be used to build it:

```cmd
go build
```

You will then get the ./go_wc executable. There are various ways to use the application. 

### Specify Flag and File

You can use the following command to utilize one of the flags above. 

```cmd
./go_wc <flag> <file>
```

This returns the value of the flag + the text file name

### Specify File

You can use the following to specify a file you would like to analyze. 

```cmd
./go_wc <file>
```

This will then create a display in the following format that showcases multiple tags. 

```cmd
  <byte count> <word count> <word count> <line count> <file name>
```

### Read from standard input

You can then also pass text from standard input, allowing you to pipe text into the command like so:

```cmd
cat <file> | ./go_wc <optional flag>
```

If you just place a flag, you will get the result as you would normally.
Alternatively, omitting the flag will get an print the byte, word and line counts on a single line, similar to sole file output. 
