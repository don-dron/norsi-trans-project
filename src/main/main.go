package main

import (
	"math/rand"
	"strconv"

	"github.com/gocql/gocql"
)

func main() {
	createTest := false
	testType := 1

	if testType == 1 {
		if createTest {
			createTestData("test_data.csv", 5000, func() []string {
				str := make([]string, 5)
				str[0] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				str[1] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				str[2] = "true"
				str[3] = "SubjectStart" + names[rand.Intn(len(names))] + names[rand.Intn(len(names))] + "SubjectEnd"
				str[4] = strconv.Itoa(rand.Intn(len(names)))
				return str
			})
		} else {
			writeData("test_data.csv", // Path to data
				"INSERT INTO test.test (dt,target,contact,direction,subject,size) VALUES( toTimeStamp(now()),?,?,?,?,?)", // Query format
				func(data [][]string) []Data { // Data builder
					result := make([]Data, 0)

					for _, array := range data {
						newItem := Data{}
						newItem.fields = make([]interface{}, 0)
						for _, item := range array {
							newItem.fields = append(newItem.fields, []byte(item))
						}
						result = append(result, newItem)
					}

					return result
				},
				func(session *gocql.Session, format string, fields []interface{}) *gocql.Query { // Query Builder
					size := len(fields)
					strs := make([]string, size)

					for j, k := range fields {
						strs[j] = string((k.([]byte))[:])
					}
					b, _ := strconv.ParseBool(strs[2])
					u, _ := strconv.ParseUint(strs[2], 10, 64)

					return session.Query(format, strs[0], strs[1], b, strs[3], u)
				})
		}
	}
}
