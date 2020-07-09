package main

import (
	"math/rand"
	"os"
	"strconv"

	"github.com/gocql/gocql"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	createTest := false
	var testType int64 = 1

	if len(os.Args) > 1 {
		testType, _ = strconv.ParseInt(os.Args[1], 10, 64)
		createTest, _ = strconv.ParseBool(os.Args[2])
	}

	if testType == 1 {
		if createTest {
			createTestData("test_data1.csv", 1000000, func() []string {
				str := make([]string, 5)
				str[0] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				str[1] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				str[2] = "true"
				str[3] = "SubjectStart" + names[rand.Intn(len(names))] + names[rand.Intn(len(names))] + "SubjectEnd"
				str[4] = strconv.Itoa(rand.Intn(len(names)))
				return str
			})
		} else {
			createTableAndKeySpace("CREATE KEYSPACE IF NOT EXISTS test WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};", "CREATE TABLE IF NOT EXISTS test.test(dt timestamp,target text,contact text,direction boolean,subject text,size int,PRIMARY KEY (target, dt, direction));")

			writeDataFromAnyFiles([]string{"test_data1.csv"}, // Path to data
				"INSERT INTO test.test (dt,target,contact,direction,subject,size) VALUES( toTimeStamp(now()),?,?,?,?,?)", // Query format
				func(data [][]string) *Queue { // Data builder
					queue := NewQueue(len(data))

					for _, array := range data {
						newItem := Data{}
						newItem.fields = make([]interface{}, 0)
						for _, item := range array {
							newItem.fields = append(newItem.fields, []byte(item))
						}
						queue.Enqueue(&newItem)
					}

					return queue
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
	} else if testType == 2 {
		createTableAndKeySpace("CREATE KEYSPACE IF NOT EXISTS test1 WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};", "CREATE TABLE IF NOT EXISTS test1.test1(dt timestamp,field0 text,field1 text,field2 text,field3 text,field4 text,field5 text,field6 text,field7 text,field8 text,field9 text,size0 int,size1 int,size2 int,size3 int,size4 int,size5 int,size6 int,size7 int,size8 int,size9 int,PRIMARY KEY (dt,field0 ,field1 ,field2 ,field3 ,field4 ,field5 ,field6 ,field7 ,field8 ,field9 ,size0 ,size1 ,size2 ,size3 ,size4 ,size5 ,size6 ,size7 ,size8 ,size9));")

		if createTest {
			createTestData("test_data2.csv", 1000000, func() []string {
				str := make([]string, 20)
				for i := 0; i < 10; i++ {
					str[i] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				}
				for i := 10; i < 20; i++ {
					str[i] = strconv.Itoa(rand.Intn(len(names)))
				}

				return str
			})
		} else {
			createTableAndKeySpace("CREATE KEYSPACE IF NOT EXISTS test1 WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};", "CREATE TABLE IF NOT EXISTS test1.test1(dt timestamp,field0 text,field1 text,field2 text,field3 text,field4 text,field5 text,field6 text,field7 text,field8 text,field9 text,size0 int,size1 int,size2 int,size3 int,size4 int,size5 int,size6 int,size7 int,size8 int,size9 int,PRIMARY KEY (dt));")

			writeDataFromAnyFiles([]string{"test_data2.csv"}, // Path to data
				"INSERT INTO test1.test1 (dt,field0 ,field1 ,field2 ,field3 ,field4 ,field5 ,field6 ,field7 ,field8 ,field9 ,size0 ,size1 ,size2 ,size3 ,size4 ,size5 ,size6 ,size7 ,size8 ,size9) VALUES( toTimeStamp(now()),? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,?)", // Query format
				func(data [][]string) *Queue { // Data builder
					queue := NewQueue(len(data))

					for _, array := range data {
						newItem := Data{}
						newItem.fields = make([]interface{}, 0)
						for _, item := range array {
							newItem.fields = append(newItem.fields, []byte(item))
						}
						queue.Enqueue(&newItem)
					}

					return queue
				},
				func(session *gocql.Session, format string, fields []interface{}) *gocql.Query { // Query Builder
					size := len(fields)
					strs := make([]string, size)

					for j, k := range fields {
						strs[j] = string((k.([]byte))[:])
					}

					u10, _ := strconv.ParseUint(strs[10], 10, 64)
					u11, _ := strconv.ParseUint(strs[11], 10, 64)
					u12, _ := strconv.ParseUint(strs[12], 10, 64)
					u13, _ := strconv.ParseUint(strs[13], 10, 64)
					u14, _ := strconv.ParseUint(strs[14], 10, 64)
					u15, _ := strconv.ParseUint(strs[15], 10, 64)
					u16, _ := strconv.ParseUint(strs[16], 10, 64)
					u17, _ := strconv.ParseUint(strs[17], 10, 64)
					u18, _ := strconv.ParseUint(strs[18], 10, 64)
					u19, _ := strconv.ParseUint(strs[19], 10, 64)

					return session.Query(format, strs[0], strs[1], strs[2], strs[3], strs[4], strs[5], strs[6], strs[7], strs[8], strs[9], u10, u11, u12, u13, u14, u15, u16, u17, u18, u19)
				})
		}
	} else if testType == 3 {
		if createTest {
			createTestData("test_data3.csv", 1000000, func() []string {
				str := make([]string, 20)
				for i := 0; i < 10; i++ {
					str[i] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				}
				for i := 10; i < 20; i++ {
					str[i] = strconv.Itoa(rand.Intn(len(names)))
				}

				return str
			})
		} else {
			createTableAndKeySpace("CREATE KEYSPACE IF NOT EXISTS test2 WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};", "CREATE TABLE IF NOT EXISTS test2.test2(dt timestamp,data blob ,PRIMARY KEY (dt,blob));")

			writeDataFromAnyFiles([]string{"test_data3.csv"}, // Path to data
				"INSERT INTO test2.test2 (dt,data) VALUES( toTimeStamp(now()),textAsBlob(?))", // Query format
				func(data [][]string) *Queue { // Data builder
					queue := NewQueue(len(data))

					for _, array := range data {
						newItem := Data{}
						newItem.fields = make([]interface{}, 0)
						for _, item := range array {
							newItem.fields = append(newItem.fields, []byte(item))
						}
						queue.Enqueue(&newItem)
					}

					return queue
				},
				func(session *gocql.Session, format string, fields []interface{}) *gocql.Query { // Query Builder
					strg := ""

					for _, k := range fields {
						strg += string((k.([]byte))[:])
					}

					return session.Query(format, strg)
				})
		}
	} else if testType == 4 {
		if createTest {
			createTestData("test_data4.csv", 1000000, func() []string {
				str := make([]string, 20)
				for i := 0; i < 10; i++ {
					str[i] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				}
				for i := 10; i < 20; i++ {
					str[i] = strconv.Itoa(rand.Intn(len(names)))
				}

				return str
			})
		} else {
			createTableAndKeySpace("CREATE KEYSPACE IF NOT EXISTS test3 WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};", "CREATE TABLE IF NOT EXISTS test3.test3(dt timestamp,data blob ,PRIMARY KEY (dt,blob));")

			writeDataFromAnyFiles([]string{"test_data4.csv"}, // Path to data
				"INSERT INTO test3.test3 (dt,data) VALUES( toTimeStamp(now()),?)", // Query format
				func(data [][]string) *Queue { // Data builder
					queue := NewQueue(len(data))

					for _, array := range data {
						msg := &ProtoTest{}

						msg.field0 = array[0]
						msg.field1 = array[1]
						msg.field2 = array[2]
						msg.field3 = array[3]
						msg.field4 = array[4]
						msg.field5 = array[5]
						msg.field6 = array[6]
						msg.field7 = array[7]
						msg.field8 = array[8]
						msg.field9 = array[9]

						u10, _ := strconv.ParseUint(array[10], 10, 64)
						u11, _ := strconv.ParseUint(array[11], 10, 64)
						u12, _ := strconv.ParseUint(array[12], 10, 64)
						u13, _ := strconv.ParseUint(array[13], 10, 64)
						u14, _ := strconv.ParseUint(array[14], 10, 64)
						u15, _ := strconv.ParseUint(array[15], 10, 64)
						u16, _ := strconv.ParseUint(array[16], 10, 64)
						u17, _ := strconv.ParseUint(array[17], 10, 64)
						u18, _ := strconv.ParseUint(array[18], 10, 64)
						u19, _ := strconv.ParseUint(array[19], 10, 64)

						msg.size0 = int32(u10)
						msg.size1 = int32(u11)
						msg.size2 = int32(u12)
						msg.size3 = int32(u13)
						msg.size4 = int32(u14)
						msg.size5 = int32(u15)
						msg.size6 = int32(u16)
						msg.size7 = int32(u17)
						msg.size8 = int32(u18)
						msg.size9 = int32(u19)

						item := Data{}
						item.fields = make([]interface{}, 1)
						item.fields[0] = msg
						queue.Enqueue(&item)
					}

					return queue
				},
				func(session *gocql.Session, format string, fields []interface{}) *gocql.Query { // Query Builder

					data, _ := proto.Marshal(fields[0].(*ProtoTest))
					return session.Query(format, data)
				})
		}
	} else if testType == 5 {
		if createTest {
			createTestData("test_data5.csv", 1000000, func() []string {
				str := make([]string, 5)
				str[0] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				str[1] = names[rand.Intn(len(names))] + strconv.Itoa(rand.Intn(len(names)))
				str[2] = "true"
				str[3] = "SubjectStart" + names[rand.Intn(len(names))] + names[rand.Intn(len(names))] + "SubjectEnd"
				str[4] = strconv.Itoa(rand.Intn(len(names)))
				return str
			})
		} else {
			createTableAndKeySpace("CREATE KEYSPACE IF NOT EXISTS test5 WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};", "CREATE TABLE IF NOT EXISTS test5.test5(dt timestamp,target text,contact text,direction boolean,subject text,size int,PRIMARY KEY (target, dt, direction,subject,size,contact));")

			writeDataFromAnyFiles([]string{"test_data5.csv"}, // Path to data
				"INSERT INTO test5.test5 (dt,target,contact,direction,subject,size) VALUES( toTimeStamp(now()),?,?,?,?,?)", // Query format
				func(data [][]string) *Queue { // Data builder
					queue := NewQueue(len(data))

					for _, array := range data {
						newItem := Data{}
						newItem.fields = make([]interface{}, 0)
						for _, item := range array {
							newItem.fields = append(newItem.fields, []byte(item))
						}
						queue.Enqueue(&newItem)
					}

					return queue
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
