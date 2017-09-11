package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"log"
)

func main() {
	filename := os.Args[1]
	if len(filename) == 0 {
		log.Fatal("No filename provided")
	}

	var stringsNumber uint64

	usedStrings := map[uint64]bool{}

	for stringsNumber == 0 || uint64(len(usedStrings)) < stringsNumber {
		var minimalString *string
		var minimalStringNumber uint64

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}

		var i uint64
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			i++
			str := scanner.Text()
			if !usedStrings[i] {
				if minimalString == nil || strings.Compare(*minimalString, str) == 1 {
					minimalString = &str
					minimalStringNumber = i
				}
			}
		}

		if i > stringsNumber {
			stringsNumber = i
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		usedStrings[minimalStringNumber] = true
		fmt.Println(*minimalString)

		file.Close()
	}

}
