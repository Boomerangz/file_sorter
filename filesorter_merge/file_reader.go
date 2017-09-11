package main

import (
	"bufio"
	"os"
)

type FileReader struct {
	file    *os.File
	scanner *bufio.Scanner

	finished bool
	str      *string
}

func (reader *FileReader) ReadString() bool {
	if reader.scanner.Scan() {
		str := reader.scanner.Text()
		reader.str = &str
		return true
	} else {
		reader.file.Close()
		os.Remove(reader.file.Name())
		reader.finished = true
		return false
	}
}
