package postgres

import (
	"context"
	"database/sql"
	"sync"
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
	log := logger.New("postgres", true)
	postgresOnce.Do(func() {
		db, err := sql.Open("postgres", source)
		if err != nil {
			log.Errorf("SINGLETON CONCURRENT DB CONNECTION ERROR !! \n", err)
			panic(err)
		}
		err = db.PingContext(context.Background())
		if err != nil {
			log.Errorf("Error pinging database: " + err.Error())
			panic(err)
		}
		log.Info("SINGLETON CONCURRENT DB CONNECTION CREATED")
		postgresClient = &Client{db}
	})
	return postgresClient
}

// ViewStats returns the status of the db.
func (c *Client) ViewStats() sql.DBStats {
	return c.Stats()
}
