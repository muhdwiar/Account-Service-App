package transaksi

import (
	"be11/project1/controllers/balance"
	"be11/project1/controllers/user"
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

func Transaksi(db *sql.DB, dataTrans entities.Transaksi, dataUser entities.User) (int, int, int, error) {
	id_receiver, err_dataReceiver := user.GetIdUser(db, dataUser)

	if err_dataReceiver != nil {
		return 0, 0, -1, err_dataReceiver
	}

	var query = "INSERT INTO transaksi (user_id, action, nominal, user_id_penerima) VALUES (?, ?, ?, ?);"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		return -1, -1, -1, errPrepare
	}

	result, errExec := statement.Exec(dataTrans.USER_ID, "Transfer", dataTrans.NOMINAL, id_receiver)

	if errExec != nil {
		return -1, -1, -1, errExec
	} else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, -1, -1, errRow
		}

		// UPDATE SALDO USER LOGIN
		UpdateSaldo_userLog := entities.Balance{}
		UpdateSaldo_userLog.ID = dataTrans.USER_ID

		balance_userLog, err_userLog := balance.Datasaldo(db, UpdateSaldo_userLog)
		if err_userLog != nil {
			return 0, -1, -1, err_userLog
		}

		UpdateSaldo_userLog.SALDO = balance_userLog - dataTrans.NOMINAL

		rBalance_userLog, errRbalance_userLog := balance.UpdateSaldo(db, UpdateSaldo_userLog)
		if errRbalance_userLog != nil {
			return 0, 0, -1, errRbalance_userLog
		}

		// UPDATE SALDO USER PENERIMA

		UpdateSaldo_receiver := entities.Balance{}
		UpdateSaldo_receiver.ID = id_receiver

		balance_receiver, err_receiver := balance.Datasaldo(db, UpdateSaldo_receiver)
		if err_receiver != nil {
			return 0, 0, -1, err_receiver
		}

		UpdateSaldo_receiver.SALDO = balance_receiver - dataTrans.NOMINAL

		rBalance_receiver, errRbalance_receiver := balance.UpdateSaldo(db, UpdateSaldo_receiver)

		if errRbalance_receiver != nil {
			return 0, 0, 0, errRbalance_receiver
		}

		return int(row), int(rBalance_userLog), int(rBalance_receiver), nil
	}
}
