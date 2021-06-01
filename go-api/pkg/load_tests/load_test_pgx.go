package event

import (
	"context"
	db "event-driven/internal/db/postgres"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func insertLoadPool(conn *pgxpool.Pool, clientN string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg := time.Now().String()
		sql := `INSERT INTO KAFKA (value) VALUES ($1)`
		//conn, err := connection.Acquire(context.Background())
		//if err != nil {
		//	panic(err)
		//} else {
		//defer conn.Release()
		_, err := conn.Exec(context.Background(), sql, msg)
		if err != nil {
			panic(err)
		} else {
			fmt.Print("\n INSERTED ", clientN, " ", msg)
		}
	}
}


func StartLoadTestPool(postgresURI string) {
	client := db.NewPostgresConnectionPool(postgresURI)
	defer client.Close()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
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
