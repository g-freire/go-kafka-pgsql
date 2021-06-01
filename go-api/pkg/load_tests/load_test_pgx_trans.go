package event

import (
	"context"
	db "event-driven/internal/db/postgres"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)


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
	}
}