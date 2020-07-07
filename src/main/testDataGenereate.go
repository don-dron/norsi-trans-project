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

func createTestData() {
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
				str := make([]string, 5)
				str[0] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				str[1] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				str[2] = "true"
				str[3] = "SubjectStart" + names[rand.Intn(len(names))] + names[rand.Intn(len(names))] + "SubjectEnd"
				str[4] = strconv.Itoa(rand.Intn(len(names)))

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
