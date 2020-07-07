package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func createTestDataHard() {
	file, err := os.Create("test_data.csv")
	rand.Seed(time.Now().UnixNano())

	if err != nil {
		fmt.Println("Unable to create file:1", err)
		os.Exit(1)
	}

	defer file.Close()

	csvWriter = csv.NewWriter(io.Writer(file))
	csvWriter.Comma = ','
	n := 5000 //0000
	w := 1000
	var wg sync.WaitGroup

	goroutines := 10

	wg.Add(goroutines)

	start := time.Now()
	for j := 0; j < goroutines; j++ {
		go func() {
			defer wg.Done()

			rows := make([][]string, 0)
			for i := 0; i < n; i++ {
				str := make([]string, 20)

				for k := 0; k < 10; k++ {
					str[k] = strconv.Itoa(rand.Intn(len(names)))
				}

				for k := 10; k < 20; k++ {
					str[k] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				}

				rows = append(rows, str)

				if (i-1)%w == 0 {
					write(rows)
					rows = make([][]string, 0)
				}
			}

			write(rows)
		}()
	}
	wg.Wait()

	elapsed := time.Now().Sub(start)
	diff := elapsed.Milliseconds()
	fmt.Print("Writting test time milliseconds ")
	fmt.Println(diff)
}
