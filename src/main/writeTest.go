package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gocql/gocql"
)

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

const (
	pageSize int = 1000000
)

func writeData(dataPath string, queryFormat string, DataBuilder func([][]string) []Data, queryBuilder func(*gocql.Session, string, []interface{}) *gocql.Query) {
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
	offset := 0

	emails := DataBuilder(ReadCSV(dataPath, offset, pageSize))

	for len(emails) != 0 {
		var ops uint64 = 0
		n := 1024
		var wg sync.WaitGroup
		start := time.Now()

		concurrecy := true // Очередь новая , на ней быстрее
		if concurrecy {
			wg.Add(n)
			queue := NewQueue(len(emails))

			for _, e := range emails {
				queue.Enqueue(&e)
			}

			start = time.Now()
			for i := 0; i < n; i++ {
				go func() {
					defer wg.Done()
					for {
						i, ok := queue.Dequeue()

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

		} else {
			wg.Add(n)
			size := len(emails)

			cellSize := size / n
			if size%n > 0 {
				cellSize++
			}

			for i := 0; i < n; i++ {
				go func(index int) {
					defer wg.Done()

					finish := (index + 1) * cellSize

					for j := index * cellSize; j < finish && j < size; j++ {
						e := emails[j]

						err := queryBuilder(session, queryFormat, e.fields).Exec()

						if err != nil {
							log.Println(err)
							return
						}

						atomic.AddUint64(&ops, 1)
						// fmt.Println(ops)
					}
				}(i)
			}
		}

		wg.Wait()
		elapsed := time.Now().Sub(start)
		diff := elapsed.Nanoseconds()
		fmt.Print(ops)
		fmt.Println(" Queries")

		fmt.Print(diff / int64(ops))
		fmt.Println(" nanoseconds time slise for one query")

		offset += pageSize
		emails = DataBuilder(ReadCSV(dataPath, offset, pageSize))
	}
}
