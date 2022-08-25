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

func GetUserBalance(db *sql.DB, dataTrans entities.Transaksi, dataUser entities.User) (entities.Balance, error) {

	var dataBalance entities.Balance
	err_select := db.QueryRow("SELECT id, saldo FROM balance WHERE id IN (SELECT id FROM user WHERE id = ? OR no_telp = ?)", dataTrans.USER_ID, dataUser.NO_TELP).Scan(&dataBalance.ID, &dataBalance.SALDO)

	if err_select != nil {
		return entities.Balance{}, err_select
	}

	return dataBalance, nil

}

func TransferBalance(db *sql.DB, dataBalance entities.Balance) (int, error) {
	var update_query = "UPDATE balance SET saldo = ? WHERE id = ?;"

	state_update, errPrep_Update := db.Prepare(update_query)

	if errPrep_Update != nil {
		return -1, errPrep_Update
	}

	result_update, errExec_Update := state_update.Exec(dataBalance.SALDO, dataBalance.ID)

	if errExec_Update != nil {
		return -1, errExec_Update
	} else {
		row, errRow_Affect := result_update.RowsAffected()

		if errRow_Affect != nil {
			return 0, errRow_Affect
		}

		return int(row), nil
	}

}
