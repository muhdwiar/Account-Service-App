package transaksi

import (
	"be11/project1/controllers/balance"
	"be11/project1/entities"
	"database/sql"
	"fmt"
	"time"
)

func InsertTrans(db *sql.DB, dataTrans entities.Transaksi) (int, error) {

	var query = "INSERT INTO transaksi (user_id, action, nominal, user_id_penerima, created_at) VALUES (?, ?, ?, ?, ?);"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		return -1, errPrepare
	}

	result, errExec := statement.Exec(dataTrans.USER_ID, dataTrans.ACTION, dataTrans.NOMINAL, dataTrans.USER_ID_PENERIMA, dataTrans.CREATED_AT)

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

func TransferBalance(db *sql.DB, trans_userLogin entities.Transaksi, userPenerima entities.User) (int, int, int, error) {

	// UPDATE BALANCE LOGIN USER
	balance_userLog, err_userLog := balance.GetUserBalance(db, trans_userLogin, entities.User{})

	if err_userLog != nil {
		return -1, -1, -1, err_userLog
	} else {
		balance_userLog.SALDO = balance_userLog.SALDO - trans_userLogin.NOMINAL
		if balance_userLog.SALDO < 0 {
			fmt.Println("GAGAL TRANSFER, SALDO ANDA KURANG DARI NOMINAL TRANSFER!!!")
			return -1, -1, -1, err_userLog
		}
	}

	row_balUser, err_balUser := balance.TransferBalance(db, balance_userLog)

	if err_balUser != nil {
		return row_balUser, -1, -1, err_balUser
	}

	// UPDATE BALANCE USER PENERIMA
	balance_penerima, err_penerima := balance.GetUserBalance(db, entities.Transaksi{}, userPenerima)

	if err_penerima != nil {
		return row_balUser, -1, -1, err_userLog
	} else {
		balance_penerima.SALDO = balance_penerima.SALDO + trans_userLogin.NOMINAL
	}

	row_balPenerima, err_balPenerima := balance.TransferBalance(db, balance_penerima)

	if err_balPenerima != nil {
		return row_balUser, row_balPenerima, -1, err_balPenerima
	}

	// INSERT DATA TRANSAKSI
	if row_balUser > 0 && row_balPenerima > 0 {

		trans_userLogin.ACTION = "Transfer"
		trans_userLogin.USER_ID_PENERIMA = balance_penerima.ID
		trans_userLogin.CREATED_AT = time.Now()

		row_transfer, err_transfer := InsertTrans(db, trans_userLogin)

		if err_transfer != nil {
			return row_balUser, row_balPenerima, row_transfer, err_transfer
		}

		fmt.Println(balance_penerima.SALDO)
		fmt.Println(balance_userLog.SALDO)

		return row_balUser, row_balPenerima, row_transfer, nil

	} else {
		return row_balUser, row_balPenerima, -1, nil
	}
}
