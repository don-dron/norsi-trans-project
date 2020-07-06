package main

import "fmt"

func main() {
	// createTestData()
	emails := CreateEmails(ReadCSV("test_data.csv"))

	fmt.Println("start writting")
	writeData(emails)
	//readData(emails)
}
