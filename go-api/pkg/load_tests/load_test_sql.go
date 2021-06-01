package event

import (
	"database/sql"
	db "event-driven/internal/db/postgres"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func insertLoad(client *sql.DB, clientN string, wg *sync.WaitGroup, stopchan chan struct{}) {
	defer wg.Done()

	for {
		select {
			default:
				msg := time.Now().String()
				sqlStatement := `INSERT INTO KAFKA (value) VALUES ($1)`
				_, err := client.Exec(sqlStatement, msg)
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

func StartLoadTest(defaultPostgresURI string) {
	var wg sync.WaitGroup
	// a channel to tell it to stop
	stopchan := make(chan struct{})

	client := db.NewPostgresSingletonClient(defaultPostgresURI)
	defer client.Close()

	start := time.Now()
	for i := 0; i < totalGoroutines; i++ {
		s := strconv.Itoa(i)
		wg.Add(1)
		//go insertLoad(client, s, &wg)
		go insertLoad(client, s, &wg, stopchan)
	}

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
