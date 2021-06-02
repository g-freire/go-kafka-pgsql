package postgres

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"
)

var (
	postgresClient *Client
	postgresOnce   sync.Once
)

type Client struct {
	Conn *sql.DB
}

func NewPostgresSingletonClient(dbHost string) *sql.DB {
	postgresOnce.Do(func() {
		db, err := sql.Open("postgres", dbHost)
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
<<<<<<< HEAD
		db.SetConnMaxLifetime(2 * time.Minute)

=======
		db.SetConnMaxLifetime(3 * time.Minute)
>>>>>>> 52b46272dcd72a17a80f4b94e8f0541c77df5e20
		if err != nil {
			log.Printf("SINGLETON CONCURRENT DB CONNECTION ERROR !! \n", err)
			panic(err)
		}
		err = db.PingContext(context.Background())
		if err != nil {
			log.Printf("Error pinging database: " + err.Error())
			panic(err)
		}
		log.Printf("SINGLETON CONCURRENT DB CONNECTION CREATED")
		postgresClient = &Client{db}
	})
	return postgresClient.Conn
}

func NewPostgresClient(source string) *sql.DB {
	db, err := sql.Open("postgres", source)
	if err != nil {
		log.Printf("DB CONNECTION ERROR !! \n", err)
		panic(err)
	}
	err = db.PingContext(context.Background())
	if err != nil {
		log.Printf("Error pinging database: " + err.Error())
		panic(err)
	}
	log.Printf("DB CONNECTION CREATED")
	postgresClient = &Client{db}
	return postgresClient.Conn
}

func (c *Client) ViewStats() sql.DBStats {
	return c.Conn.Stats()
}
