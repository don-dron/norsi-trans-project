package main

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

func readData() {
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

	err = session.Query("CREATE TABLE IF NOT EXISTS test.test (dt timestamp,target text,contact text,direction boolean,subject text,size bigint,PRIMARY KEY ((target), dt, direction));").Exec()

	if err != nil {
		log.Println(err)
		return
	}
}
