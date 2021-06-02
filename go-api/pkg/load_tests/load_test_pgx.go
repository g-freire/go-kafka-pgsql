package event

import (
	"context"
	db "event-driven/internal/db/postgres"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func insertLoadPool(conn *pgxpool.Pool, clientN string, wg *sync.WaitGroup, stopchan chan struct{}) {
	defer wg.Done()
	for {
		select {
		default:
			msg := time.Now().String()
			sql := `INSERT INTO KAFKA (value) VALUES ($1)`
			_, err := conn.Exec(context.Background(), sql, msg)
			if err != nil {
				panic(err)
			} else {
				fmt.Print("\n INSERTED ", clientN, " ", msg)
			}
		case <-stopchan:
			return
		}
	}
}

func StartLoadTestPool(postgresURI string) {
	var wg sync.WaitGroup
	// a channel to tell it to stop
	stopchan := make(chan struct{})

	client := db.NewPostgresConnectionPool(postgresURI)
	defer client.Close()

	start := time.Now()
	for i := 0; i < totalGoroutines; i++ {
		s := strconv.Itoa(i)
		wg.Add(1)
		go insertLoadPool(client, s, &wg, stopchan)
		//time.Sleep(time.Duration(1) * time.Second)
	}
	//wg.Wait()

	// sends stop signal
	time.Sleep(testTimeSeconds * time.Second)
	close(stopchan) // tell it to stop
	<-stopchan      // wait for it to have stopped
	log.Println("Stopped.")

	// will panic
	//wg.Done()

	elapsed := time.Since(start)
	log.Printf("\n ------------------------------------")
	log.Printf("Num. goroutines", totalGoroutines )
	log.Printf("Test time %v", testTimeSeconds , " seconds")
	log.Printf("Process took %s", elapsed )
	log.Printf("------------------------------------ \n")
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
