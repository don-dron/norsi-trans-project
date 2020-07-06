package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gocql/gocql"
)

func writeData(emails []Email) {
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

	// var ops uint64 = 0
	n := 1024
	var wg sync.WaitGroup
	errs := make(chan error)
	start := time.Now()

	if concurrecy {
		wg.Add(n)
		queue := NewQueue()

		for _, e := range emails {
			queue.Enqueue(&e)
		}

		start = time.Now()
		query := "INSERT INTO test.test (dt,target,contact,direction,subject,size) VALUES( toTimeStamp(now()),?,?,?,?,?)"
		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
				for {
					e, ok := queue.Dequeue().(*Email)

					if !ok {
						return
					}

					err := session.Query(query, e.target, e.contact, e.direction, e.subject, e.size).Exec()

					if err != nil {
						errs <- err
						return
					}
					// atomic.AddUint64(&ops, 1)
					// fmt.Println(ops)
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
		for i := 0; i < n-1; i++ {
			go func(index int) {
				defer wg.Done()

				finish := (index + 1) * cellSize

				for j := index * cellSize; j < finish; j++ {
					e := emails[j]

					err := session.Query(query, e.target, e.contact, e.direction, e.subject, e.size).Exec()

					if err != nil {
						errs <- err
					}

					// atomic.AddUint64(&ops, 1)
					// fmt.Println(ops)
				}
			}(i)
		}

		go func(index int) {
			defer wg.Done()

			finish := (index + 1) * cellSize

			for j := index * cellSize; j < finish && j < size; j++ {
				e := emails[j]

				err := session.Query(query, e.target, e.contact, e.direction, e.subject, e.size).Exec()

				if err != nil {
					errs <- err
				}
				// atomic.AddUint64(&ops, 1)
				// fmt.Println(ops)
			}
		}(n - 1)
	}

	wg.Wait()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed.Milliseconds())
	close(errs)

	for err := range errs {
		if err != nil {
			log.Println(err)
		}
	}
}
