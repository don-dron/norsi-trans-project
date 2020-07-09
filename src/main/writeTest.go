package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gocql/gocql"
)

// QueryBuilder Конструирует запрос к бд на основе сессии,формата запроса и данных
type QueryBuilder func(*gocql.Session, string, []interface{}) *gocql.Query

// DataBuilder конструирует модель данных из прочитанных в файле строк
type DataBuilder func([][]string) *Queue

const (
	// Чтение файлов по кускам очень плохо, в идеале нужно читать один раз , если файл большой, больше 10 млн строк, то читаем дальше , пока не дочитаем
	pageSize int = 1000000
)

func writeDataFromAnyFiles(dataPaths []string, queryFormat string, dataBuilder DataBuilder, queryBuilder QueryBuilder) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	var wg sync.WaitGroup
	wg.Add(len(dataPaths))
	for _, path := range dataPaths {
		go func() {
			defer wg.Done()
			readCsvAndLoad(session, queryFormat, path, dataBuilder, queryBuilder)
		}()
	}
	wg.Wait()
}

func writeData(session *gocql.Session, dataPath string, queryFormat string, dataBuilder DataBuilder, queryBuilder QueryBuilder) {
	readCsvAndLoad(session, queryFormat, dataPath, dataBuilder, queryBuilder)
}

func readCsvAndLoad(session *gocql.Session, queryFormat string, dataPath string, dataBuilder DataBuilder, queryBuilder QueryBuilder) {
	offset := 0
	emails := dataBuilder(ReadCSV(dataPath, offset, pageSize))
	runtime.GOMAXPROCS(runtime.NumCPU())
	for emails.Size() != 0 {
		var ops uint64 = 0
		n := 1000
		var wg sync.WaitGroup
		start := time.Now()

		wg.Add(n)

		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
				for {
					i, ok := emails.Dequeue()

					if !ok {
						return
					}
					e := i.(*Data)

					err := queryBuilder(session, queryFormat, e.fields).Exec()

					if err != nil {
						log.Println(err)
						return
					}
					atomic.AddUint64(&ops, 1)
					// fmt.Println(ops)
				}
			}()
		}

		wg.Wait()
		elapsed := time.Now().Sub(start)
		diff := elapsed.Nanoseconds()
		fmt.Print(ops)
		fmt.Println(" Queries")

		fmt.Print(elapsed.Milliseconds())
		fmt.Println(" All time")

		fmt.Print(diff / int64(ops))
		fmt.Println(" nanoseconds time slise for one query")

		offset += pageSize
		emails = dataBuilder(ReadCSV(dataPath, offset, pageSize))
	}
}

func createTableAndKeySpace(keySpaceFormat string, tableFormat string) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	err = session.Query(keySpaceFormat).Exec()
	if err != nil {
		log.Println(err)
		return
	}

	err = session.Query(tableFormat).Exec()

	if err != nil {
		log.Println(err)
		return
	}
}
