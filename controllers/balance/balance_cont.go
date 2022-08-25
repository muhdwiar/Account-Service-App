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

func UpdateSaldo(db *sql.DB, Balance entities.Balance) (int, error) {
	var query = "update balance set saldo = ? where id = ?"
	stat, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		return -1, errPrepare
	}

	result, errExec := stat.Exec(Balance.SALDO, Balance.ID)

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

func Datasaldo(db *sql.DB, balance entities.Balance) (int, error) {
	var Balance = entities.Balance{}

	err := db.QueryRow("SELECT saldo FROM balance where id = ?", balance.ID).
		Scan(&Balance.SALDO)

	if err != nil {
		return Balance.SALDO, err
	}
	return Balance.SALDO, nil
}
