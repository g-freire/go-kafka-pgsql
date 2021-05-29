package event

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	postgresClient *Client
	postgresOnce   sync.Once
)

// Client is a client for the PostgreSQL db engine.
type Client struct {
	*sql.DB
}

// NewPostgresClient returns a new client for postgres.
func NewPostgresClient(source string) *Client {
	db, err := sql.Open("postgres", source)
	if err != nil {
		fmt.Print("SINGLETON CONCURRENT DB CONNECTION ERR !! \n", err)
		panic(err)
	}
	err = db.PingContext(context.Background())
	if err != nil {
		fmt.Print("ERR pinging database: " + err.Error())
		panic(err)
	}
	fmt.Print("SINGLETON CONCURRENT DB CONNECTION CREATED")
	postgresClient = &Client{db}
	return postgresClient
}

// ViewStats returns the status of the db.
func (c *Client) ViewStats() sql.DBStats {
	return c.Stats()
}

const defaultPostgresURI = "postgres://admin:admin@localhost:6543/admin?sslmode=disable"

func insertLoad(client *Client, clientN string) {

	for {
		msg := time.Now().String()

		// pgsql insert that has injection protection and returns the result
		// sqlStatement := `INSERT INTO KAFKA(values)
		// 				    VALUES ($1)
		// 				    RETURNING *;`

		sqlStatement := `INSERT INTO KAFKA (value) VALUES ($1)`

		_, err := client.Exec(sqlStatement, msg)
		if err != nil {
			panic(err)
		} else {
			fmt.Print("\n INSERTED ", clientN, " ", msg)
		}
		// time.Sleep(time.Second)
	}
}

func StartLoadTest() {
	postgresURI := os.Getenv("DATABASE_URI")
	if postgresURI == "" {
		postgresURI = defaultPostgresURI
	}
	// load test the db session limit
	client := NewPostgresClient(postgresURI)
	var i = 0
	for {
		defer client.Close()
		s := strconv.Itoa(i)
		go insertLoad(client, s)
		i++
	}
}

func StartLoadTestSessions() {
	postgresURI := os.Getenv("DATABASE_URI")
	if postgresURI == "" {
		postgresURI = defaultPostgresURI
	}

	//load test the db session limit
	for i := 0; i < 5000; i++ {
		client := NewPostgresClient(postgresURI)
		defer client.Close()
		s := strconv.Itoa(i)
		go insertLoad(client, s)
	}
}
