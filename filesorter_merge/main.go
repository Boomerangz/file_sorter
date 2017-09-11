package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"log"
)

var MAX_BUFFER_SIZE int = 100000

func main() {
	filename := os.Args[1]
	if len(filename) == 0 {
		log.Fatal("No filename provided")
	}

	countOfFiles := DivideFile(filename, "/tmp/sorted_file_%d", MAX_BUFFER_SIZE)
	readersList := CreateFileReaders(countOfFiles, "/tmp/sorted_file_%d")

	for {
		var minimalReader *FileReader
		for _, reader := range readersList {
			if !reader.finished && (minimalReader == nil || strings.Compare(*reader.str, *minimalReader.str) == -1) {
				minimalReader = reader
			}
		}
		if minimalReader == nil {
			break
		}

		fmt.Println(*minimalReader.str)
		minimalReader.ReadString()
	}
}

func DivideFile(filename string, fileTemplate string, bufferSize int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	stringsBuffer := StringsBuffer{}

	scanner := bufio.NewScanner(file)
	countOfWrittenFiles := 0
	for scanner.Scan() {
		str := scanner.Text()
		stringsBuffer = append(stringsBuffer, str)
		if len(stringsBuffer) >= bufferSize {
			countOfWrittenFiles++
			SortAndWrite(&stringsBuffer, countOfWrittenFiles, fileTemplate)
			stringsBuffer = StringsBuffer{}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()

	if stringsBuffer.Len() > 0 {
		countOfWrittenFiles++
		SortAndWrite(&stringsBuffer, countOfWrittenFiles, fileTemplate)
		stringsBuffer = StringsBuffer{}
	}
	return countOfWrittenFiles
}

func SortAndWrite(buffer *StringsBuffer, n int, fileTemplate string) {
	sort.Sort(buffer)
	filename := fmt.Sprintf(fileTemplate, n)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range *buffer {
		file.WriteString(s)
		file.WriteString("\n")
	}
}

func CreateFileReaders(countOfFiles int, fileTemplate string) []*FileReader {
	readersList := make([]*FileReader, countOfFiles)
	for i := 1; i <= countOfFiles; i++ {
		file, err := os.Open(fmt.Sprintf(fileTemplate, i))
		if err != nil {
			log.Fatal(err)
		}

		reader := FileReader{file: file, scanner: bufio.NewScanner(file), finished: false}
		readersList[i-1] = &reader
		reader.ReadString()
	}
	return readersList
}
