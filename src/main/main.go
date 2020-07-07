package main

func main() {
	createTest := false
	if createTest {
		createTestData()
	} else {
		writeData()
	}
}
