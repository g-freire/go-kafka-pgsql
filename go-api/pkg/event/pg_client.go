package event

import (
	"context"
	"database/sql"
	db "event-driven/internal/db/postgres"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)


func insertLoad(client *sql.DB, clientN string) {
	for{
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

func insertLoadPool(connection *pgxpool.Pool, clientN string) {
	for{
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

	for i := 0; i<10 ; i++{
		s := strconv.Itoa(i)
		go insertLoad(client, s)
		go insertLoad(client, s)
		go insertLoad(client, s)

		//time.Sleep(time.Duration(1) * time.Second)
	}
	time.Sleep(time.Duration(1) * time.Second)

}

func StartLoadTestPool(postgresURI string) {
	client := db.NewPostgresConnectionPool(postgresURI)
	defer client.Close()
	for i := 0; ; i++{
		s := strconv.Itoa(i)
		go insertLoadPool(client, s)
		//time.Sleep(time.Duration(1) * time.Second)
	}
}

// load test the db session with many inserts


func StartLoadTestSessions(defaultPostgresURI string) {
	postgresURI := os.Getenv("DATABASE_URI")
	if postgresURI == "" {
		postgresURI = defaultPostgresURI
	}
	client := db.NewPostgresConnectionPool(postgresURI)
	defer client.Close()
	//load test the db session limit
	for i := 0; i < 5000; i++ {
		s := strconv.Itoa(i)
		insertLoadPool(client, s)
	}
}
