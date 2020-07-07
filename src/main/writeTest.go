package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gocql/gocql"
)

func writeData() {
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

	err = session.Query("CREATE KEYSPACE IF NOT EXISTS test WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};").Exec()
	if err != nil {
		log.Println(err)
		return
	}

	err = session.Query("CREATE TABLE IF NOT EXISTS test.test (dt timestamp,target text,contact text,direction boolean,subject text,size int,PRIMARY KEY (target, dt, direction));").Exec()

	if err != nil {
		log.Println(err)
		return
	}

	//
	//
	// Возможно стоит переписать на параллельное обращение к массиву , а не использовать лок-фри очередь,
	// 1)Не будет долбежки на касах внутри очереди(синхронизации на очереди)
	// 2)Из-за обращения к массиву не будет инвалидации кэшей
	// 3)Не надо будет перегонять массив в очередь
	// 4)Чтение по указателю дороже - происходит обращение в память, если читать массив на прямую такого нет, будут вступать в роль кэши
	//
	// С массивом быстрее на 2 секунды на 1 000 000 данных
	//

	concurrecy := false

	var ops uint64 = 0

	offset := 0
	page_size := 1000000

	emails := CreateEmails(ReadCSV("test_data.csv", offset, page_size))

	for len(emails) != 0 {
		n := 1024
		var wg sync.WaitGroup
		errs := make(chan error)
		start := time.Now()

		if concurrecy {
			wg.Add(n)
			queue := NewQueue(len(emails))

			for _, e := range emails {
				queue.Enqueue(&e)
			}

			start = time.Now()
			query := "INSERT INTO test.test (dt,target,contact,direction,subject,size) VALUES( toTimeStamp(now()),?,?,?,?,?)"
			for i := 0; i < n; i++ {
				go func() {
					defer wg.Done()
					for {
						i, ok := queue.Dequeue()

						if !ok {
							return
						}
						e := i.(*Email)
						err := session.Query(query, e.target, e.contact, e.direction, e.subject, e.size).Exec()

						if err != nil {
							errs <- err
							return
						}
						atomic.AddUint64(&ops, 1)
					}
				}()
			}

		} else {
			n = 1024
			wg.Add(n)
			size := len(emails)

			cellSize := size / n
			if size%n > 0 {
				cellSize++
			}
			start = time.Now()
			query := "INSERT INTO test.test (dt,target,contact,direction,subject,size) VALUES( toTimeStamp(now()),?,?,?,?,?)"
			for i := 0; i < n; i++ {
				go func(index int) {
					defer wg.Done()

					finish := (index + 1) * cellSize

					for j := index * cellSize; j < finish && j < size; j++ {
						e := emails[j]

						err := session.Query(query, e.target, e.contact, e.direction, e.subject, e.size).Exec()

						if err != nil {
							log.Println(err)
							errs <- err
						}

						atomic.AddUint64(&ops, 1)
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
		close(errs)
		for err := range errs {
			if err != nil {
				log.Println(err)
			}
		}
		offset += page_size
		emails = CreateEmails(ReadCSV("test_data.csv", offset, page_size))
	}
}
