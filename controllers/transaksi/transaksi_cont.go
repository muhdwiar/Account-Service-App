package transaksi

import (
	"be11/project1/controllers/balance"
	"be11/project1/entities"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func TopUp(db *sql.DB, datransaksi entities.Transaksi, datauser entities.User) (int, int, error) {
	var query = "insert into transaksi (user_id, action, nominal, user_id_penerima) values (?, ?, ?, ?)"
	stat, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		return -1, -1, errPrepare
	}

	result, errExec := stat.Exec(datauser.ID, "Top UP", datransaksi.NOMINAL, datauser.ID)

	if errExec != nil {
		return -1, -1, errExec
	} else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, -1, errRow
		}
		UpdateSaldo := entities.Balance{}
		UpdateSaldo.ID = datauser.ID

		temp, tempErr := balance.Datasaldo(db, UpdateSaldo)
		if tempErr != nil {
			return 0, -1, tempErr
		}

		UpdateSaldo.SALDO = temp + datransaksi.NOMINAL

		rBalance, errRbalance := balance.UpdateSaldo(db, UpdateSaldo)
		if errRbalance != nil {
			return 0, 0, errRbalance
		}
		return int(row), int(rBalance), nil
	}

}
