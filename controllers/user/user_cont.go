package user

import (
	"be11/project1/entities"
	"database/sql"
)

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
