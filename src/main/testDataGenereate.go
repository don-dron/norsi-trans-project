package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

var names []string = []string{
	"Liam", "Emma",
	"Noah", "Olivia",
	"Mason", "Ava",
	"Ethan", "Sophia",
	"Logan", "Isabella",
	"Lucas", "Mia",
	"Jackson", "Charlotte",
	"Aiden", "Amelia",
	"Oliver", "Emily",
	"Jacob", "Madison",
	"Elijah", "Harper",
	"Alexander", "Abigail",
	"James", "Avery",
	"Benjamin", "Lily",
	"Jack", "Ella",
	"Luke", "Chloe",
	"William", "Evelyn",
	"Michael", "Sofia",
	"Owen", "Aria",
	"Daniel", "Ellie",
	"Carter", "Aubrey",
	"Gabriel", "Scarlett",
	"Henry", "Zoey",
	"Matthew", "Hannah",
	"Wyatt", "Audrey",
	"Caleb", "Grace",
	"Jayden", "Addison",
	"Nathan", "Zoe",
	"Ryan", "Elizabeth",
	"Isaac", "Nora"}

var mu sync.Mutex
var csvWriter *csv.Writer

func createTestData(dataPath string, dataCount int, dataCreator func() []string) {
	file, err := os.Create(dataPath)
	rand.Seed(time.Now().UnixNano())

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()

	csvWriter = csv.NewWriter(io.Writer(file))
	csvWriter.Comma = ','

	w := 1000
	var wg sync.WaitGroup

	goroutines := 10
	n := dataCount

	wg.Add(goroutines)

	start := time.Now()
	for j := 0; j < goroutines; j++ {
		go func() {
			defer wg.Done()

			rows := make([][]string, 0)
			for i := 0; i < n; i++ {
				str := dataCreator()

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

func write(strs [][]string) {
	mu.Lock()
	defer mu.Unlock()
	csvWriter.WriteAll(strs)
}
