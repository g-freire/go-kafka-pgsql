package event

import (
	"database/sql"
	db "event-driven/internal/db/postgres"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)


func insertLoad(client *sql.DB, clientN string) {
		msg := time.Now().String()
		sqlStatement := `INSERT INTO KAFKA (value) VALUES ($1)`
		_, err := client.Exec(sqlStatement, msg)
		if err != nil {
			panic(err)
		} else {
			fmt.Print("\n INSERTED ", clientN, " ", msg)
		}

}

func StartLoadTest(defaultPostgresURI string) {
	postgresURI := os.Getenv("DATABASE_URI")
	if postgresURI == "" {
		postgresURI = defaultPostgresURI
	}

	// load test the db with many inserts
	client := db.NewPostgresSingletonClient(postgresURI)
	defer client.Close()
	for i := 0; ; i++{
		s := strconv.Itoa(i)
		insertLoad(client, s)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func StartLoadTestSessions(defaultPostgresURI string) {
	postgresURI := os.Getenv("DATABASE_URI")
	if postgresURI == "" {
		postgresURI = defaultPostgresURI
	}

	//load test the db session limit
	for i := 0; i < 5000; i++ {
		client := db.NewPostgresClient(postgresURI)
		defer client.Close()
		s := strconv.Itoa(i)
		go insertLoad(client, s)
	}
}
