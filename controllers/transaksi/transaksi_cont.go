package transaksi

import (
	"be11/project1/entities"
	"database/sql"
)

func GetTransfer_asSender(db *sql.DB, dataUser entities.User) ([]entities.Transaksi, error) {
	action := "Transfer"
	result_select, err_select := db.Query("SELECT tr.id, tr.nominal, tr.created_at, u.nama, u.no_telp FROM transaksi tr INNER JOIN user u ON tr.user_id_penerima = u.id WHERE tr.user_id = ? AND tr.action = ? ORDER BY tr.created_at;", &dataUser.ID, &action)

	if err_select != nil {
		return nil, err_select
	}

	var dataTransferAll []entities.Transaksi

	for result_select.Next() {
		var row_trans entities.Transaksi

		errScan := result_select.Scan(&row_trans.ID, &row_trans.NOMINAL, &row_trans.CREATED_AT, &row_trans.USER.NAMA, &row_trans.USER.NO_TELP)

		if errScan != nil {
			return nil, errScan
		} else {
			dataTransferAll = append(dataTransferAll, row_trans)

		}
	}

	return dataTransferAll, nil

}

func GetTransfer_asReceiver(db *sql.DB, dataUser entities.User) ([]entities.Transaksi, error) {

	result_select, err_select := db.Query("SELECT tr.id, tr.nominal, tr.created_at, u.nama, u.no_telp FROM transaksi tr INNER JOIN user u ON tr.user_id = u.id WHERE tr.user_id_penerima = ? AND tr.action = ? ORDER BY tr.created_at;", &dataUser.ID, "Transfer")

	if err_select != nil {
		return nil, err_select
	}

	var dataTransferAll []entities.Transaksi

	for result_select.Next() {
		var row_trans entities.Transaksi

		errScan := result_select.Scan(&row_trans.ID, &row_trans.NOMINAL, &row_trans.CREATED_AT, &row_trans.USER.NAMA, &row_trans.USER.NO_TELP)

		if errScan != nil {
			return nil, errScan
		} else {
			dataTransferAll = append(dataTransferAll, row_trans)

		}
	}

	return dataTransferAll, nil

}
