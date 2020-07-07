package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

func ReadCSV(path string, offset int, stringCount int) [][]string {
	result := make([][]string, 0)

	csvFile, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(csvFile)
	r.Comma = ','

	start := time.Now()
	n := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		n++

		if n > offset+stringCount {
			break
		}
		if n > offset {
			result = append(result, record)
		}
	}

	elapsed := time.Now().Sub(start)
	diff := elapsed.Milliseconds()
	fmt.Print("Reader time ")
	fmt.Println(diff)

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
