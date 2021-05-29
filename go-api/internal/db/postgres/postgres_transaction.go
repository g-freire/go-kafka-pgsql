package postgres

import (
	"database/sql"
	"log"
)

func RollbackTx(tx *sql.Tx, err error) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		log.Printf("\n[ERROR]: UNABLE TO ROLLBACK \n", rollbackErr)
	}
	log.Printf("\n[ERROR]: TRANSACTION COULD NOT EXEC CONTEXT \n", err)
}
