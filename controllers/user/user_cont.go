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

func ReadProfile(db *sql.DB, datauser entities.User) (entities.User, error) {

	dataUser := entities.User{}
	err := db.QueryRow("SELECT u.id, u.nama, u.no_telp, u.created_at, u.updated_at, b.saldo FROM user u INNER JOIN balance b ON u.id = b.id WHERE u.id = ?", datauser.ID).
		Scan(&dataUser.ID, &dataUser.NAMA, &dataUser.NO_TELP, &dataUser.CREATED_AT, &dataUser.UPDATED_AT, &dataUser.BALANCE.SALDO)
	if err != nil {
		return dataUser, err
	}
	return dataUser, nil
}

func UpdateProfile(db *sql.DB, datauser entities.User, user entities.User) (int, error) {
	// _, err := db.Exec("update user set nama = ?, no_telp = ?, password = ? where id = ?", datauser.ID)

	var query = "update user set nama = ?, no_telp = ?, password = ? where id = ?"

	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return -1, errPrepare
	}
	result, errExec := statement.Exec(user.NAMA, user.NO_TELP, user.PASSWORD, datauser.ID)
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

		return int(row_user), nil
	}
}

func GetIdUser(db *sql.DB, datauser entities.User) (int, error) {

	dataUser := entities.User{}
	err := db.QueryRow("SELECT id FROM user WHERE no_telp = ?", datauser.NO_TELP).Scan(&dataUser.ID)
	if err != nil {
		return dataUser.ID, err
	}
	return dataUser.ID, nil
}
