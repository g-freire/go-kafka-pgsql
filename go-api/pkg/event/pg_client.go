package event

import (
	"context"
	"database/sql"
	db "event-driven/internal/db/postgres"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)


func insertLoad(client *sql.DB, clientN string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg := time.Now().String()
		sqlStatement := `INSERT INTO KAFKA (value) VALUES ($1)`
		_, err := client.Exec(sqlStatement, msg)
		if err != nil {
			panic(err)
		} else {
			fmt.Print("\n INSERTED ", clientN, " ", msg)
		}
	}
}

func insertLoadPool(connection *pgxpool.Pool, clientN string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg := time.Now().String()
		sqlStatement := `INSERT INTO KAFKA (value) VALUES ($1)`
		//
		//conn, err := connection.Acquire(context.Background())
		//if err != nil {
		//	panic(err)
		//} else {
		//defer conn.Release()
		_, err := connection.Exec(context.Background(), sqlStatement, msg)
		if err != nil {
			panic(err)
		} else {
			fmt.Print("\n INSERTED ", clientN, " ", msg)
		}
	}
}

//_, err = connection.Exec(context.Background(), sqlStatement, msg)
//if err != nil {
//	panic(err)
//} else {
//	fmt.Print("\n INSERTED ", clientN, " ", msg)
//}
// load test the db with many inserts

func StartLoadTest(defaultPostgresURI string) {
	client := db.NewPostgresSingletonClient(defaultPostgresURI)
	defer client.Close()

	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		s := strconv.Itoa(i)
		wg.Add(1)
		//go insertLoad(client, s, &wg)
		go insertLoad(client, s, &wg)
	}
	wg.Wait()

}

func StartLoadTestPool(postgresURI string) {
	client := db.NewPostgresConnectionPool(postgresURI)
	defer client.Close()
	var wg sync.WaitGroup
	for i := 0; i < 500000; i++ {
		s := strconv.Itoa(i)
		wg.Add(1)
		go insertLoadPool(client, s, &wg)
		//time.Sleep(time.Duration(1) * time.Second)
	}
	wg.Wait()
}

// load test the db session with many inserts

//func StartLoadTestSessions(defaultPostgresURI string) {
//	postgresURI := os.Getenv("DATABASE_URI")
//	if postgresURI == "" {
//		postgresURI = defaultPostgresURI
//	}
//	client := db.NewPostgresConnectionPool(postgresURI)
//	defer client.Close()
//	//load test the db session limit
//	for i := 0; i < 5000; i++ {
//		s := strconv.Itoa(i)
//		insertLoadPool(client, s)
//	}
//}
