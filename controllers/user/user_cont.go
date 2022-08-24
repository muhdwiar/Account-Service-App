package user

import (
	"be11/project1/controllers/balance"
	"be11/project1/entities"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Registrasi(db *sql.DB, newUser entities.User) (int, int, error) {

	var query = "insert into user (nama, no_telp, password) values (?, ?, ?)"
	statement, errPrepare := db.Prepare(query)

	if errPrepare != nil {
		return -1, -1, errPrepare
	}

	result, errExec := statement.Exec(newUser.NAMA, newUser.NO_TELP, newUser.PASSWORD)

	if errExec != nil {
		return -1, -1, errExec
	} else {
		row, errRow := result.RowsAffected()
		if errRow != nil {
			return 0, -1, errRow
		}
		id, err := result.LastInsertId()
		if err != nil {
			return 0, -1, err
		}
		newBalance := entities.Balance{}
		newBalance.ID = int(id)
		newBalance.SALDO = 0

		row_balance, errRowb := balance.InputBalance(db, newBalance)
		if errRowb != nil {
			return 0, 0, errRowb
		}
		return int(row), int(row_balance), nil
	}
}

func LoginUser(db *sql.DB, loginUser entities.User) (entities.User, error) {
	result, err := db.Query("SELECT id, nama, no_telp FROM user WHERE no_telp = ? AND password = ?", &loginUser.NO_TELP, &loginUser.PASSWORD)

	if err != nil {
		return entities.User{}, err
	}

	var dataUserLogin entities.User

	for result.Next() {

		errScan := result.Scan(&dataUserLogin.ID, &dataUserLogin.NAMA, &dataUserLogin.NO_TELP)

		if errScan != nil {
			return entities.User{}, errScan
		}
	}

	return dataUserLogin, nil

}

func DeleteUser(db *sql.DB, deleteUser entities.User) (int, error) {
	var delete_query = "DELETE FROM user WHERE id = ?"

	statement, errPrep := db.Prepare(delete_query)

	if errPrep != nil {
		return -1, errPrep
	}

	result, errExec := statement.Exec(deleteUser.ID)

	if errExec != nil {
		return -1, errExec
	} else {
		row_user, errRow_Affect := result.RowsAffected()

		if errRow_Affect != nil {
			return 0, errRow_Affect
		}

		deleteBalance := entities.Balance{}
		deleteBalance.ID = deleteUser.ID

		row_balance, err := balance.DeleteBalance(db, deleteBalance)

		if err != nil {
			return 1, 0, err
		}

		return int(row_user), nil
	}
}
