package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	var stringsCountStr string
	var stringsSizeStr string

	flag.StringVar(&stringsCountStr, "strcount", "1024", "count of generated strings")
	flag.StringVar(&stringsSizeStr, "strsize", "255", "size of generated strings")

	flag.Parse()

	var stringsCount int
	var stringsSize int

	stringsCount, err := strconv.Atoi(stringsCountStr)
	if err != nil || stringsCount <= 0 {
		log.Fatal("strcount must be positive natural number")
	}

	stringsSize, err = strconv.Atoi(stringsSizeStr)
	if err != nil || stringsSize <= 0 {
		log.Fatal("stringsSize must be positive natural number")
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < stringsCount; i++ {
		fmt.Println(RandString(stringsSize))
	}

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, rand.Intn(n)+1)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
