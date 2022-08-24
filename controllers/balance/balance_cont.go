package balance

import (
	"be11/project1/entities"
	"database/sql"
)

func InputBalance(db *sql.DB, newBalance entities.Balance) (int, error) {

	var query = "insert into balance (id, saldo) values (?, ?)"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		return -1, errPrepare
	}

	result, errExec := statement.Exec(newBalance.ID, newBalance.SALDO)

	if errExec != nil {
		return -1, errExec
	} else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, errRow
		}
		return int(row), nil
	}
}
