package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

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
