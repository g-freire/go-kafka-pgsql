package postgres

import (
	"database/sql"
)

func RollbackTx(tx *sql.Tx, log logger.Logger, err error) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		log.Errorf("\n[ERROR]: UNABLE TO ROLLBACK \n", rollbackErr)
	}
	log.Errorf("\n[ERROR]: TRANSACTION COULD NOT EXEC CONTEXT \n", err)
}
