package postgres

import (
	"context"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var (
	postgresClientPool *ClientPool
)

type ClientPool struct {
	Conn *pgxpool.Pool
}

func NewPostgresConnectionPool(dbHost string) *pgxpool.Pool {
	postgresOnce.Do(func() {
		pool, err := pgxpool.Connect(context.Background(), dbHost)
		if err != nil {
			log.Fatalf("Couldn't connect to the database. Reason %v", err)
		}
		//pool.Stat()
		postgresClientPool = &ClientPool{Conn: pool}
	})
    log.Printf("DB CONNECTION POOL CREATED")
    return postgresClientPool.Conn
}