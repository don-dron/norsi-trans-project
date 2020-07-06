package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

func createTestData() {
	file, err := os.Create("test_data.csv")
	rand.Seed(time.Now().UnixNano())

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()

	for i := 0; i < 1000000; i++ {
		file.WriteString(names[rand.Intn(len(names))] + " " + names[rand.Intn(len(names))] + " SubjectStart" + names[rand.Intn(len(names))] + names[rand.Intn(len(names))] + "SubjectEnd " + strconv.Itoa(rand.Intn(len(names))) + "\n")
	}
}
