package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Email struct {
	timestamp time.Time
	target    string
	contact   string
	direction bool
	subject   string
	size      uint64
}

func ReadCSV(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	stat, _ := file.Stat()

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	result := make([][]string, 0)

	data := string(bs)

	strArray := strings.Split(data, "\n")

	for _, s := range strArray {
		result = append(result, strings.Split(s, " "))
	}

	return result
}

func CreateEmails(data [][]string) []Email {
	result := make([]Email, 0)

	for _, array := range data {

		ui, _ := strconv.ParseUint(array[3], 10, 64)

		result = append(result, Email{time.Now(), array[0], array[1], true, array[2], ui})
		result = append(result, Email{time.Now(), array[1], array[0], false, array[2], ui})
	}

	return result
}
