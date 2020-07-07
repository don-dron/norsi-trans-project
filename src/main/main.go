package main

import (
	"strconv"

	"github.com/gocql/gocql"
)

type Data struct {
	fields []interface{}
}

func main() {
	createTest := false
	if createTest {
		createTestData()
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
