package event

import (
	"context"
	"database/sql"
	db "event-driven/internal/db/postgres"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
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



func insertLoadPoolTran(connection *pgxpool.Pool, clientN string, wg *sync.WaitGroup) {
	defer wg.Done()
	var ctx = context.TODO()
	for {
		msg := time.Now().String()
		tx, err := connection.BeginTx(context.TODO(), pgx.TxOptions{IsoLevel: "serializable"})
		if err != nil {
			log.Printf("\n[ERROR]: TRANSACTION COULD NOT BEGIN\n", err)
		}
		defer tx.Rollback(context.TODO())

		tsql := `INSERT INTO KAFKA (value) VALUES ($1)`
		tran, err := tx.Exec(ctx, tsql, msg)
		if err != nil {
			db.RollbackTxPgx(tx, err)
			return
		}

		rowsAffected := tran.RowsAffected()
		if err != nil || rowsAffected != 1 {
			log.Print(err)
			return
		}
		err = tx.Commit(context.TODO())
		if err != nil {
			log.Printf("\n[ERROR]: TRANSACTION COULD NOT COMMIT \n", err)
		}else{
			fmt.Print("\n INSERT COMMITED  ", clientN, " ", msg)
		}


		//conn, err := connection.Acquire(context.Background())
		//if err != nil {
		//	panic(err)
		//} else {
		//defer conn.Release()
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
