package main

import (
	"be11/project1/config"
	"be11/project1/entities"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//  go get github.com/joho/godotenv
//  go get -u github.com/go-sql-driver/mysql

func main() {

	err_env := godotenv.Load()

	if err_env != nil {
		fmt.Println("Error load env file")
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	fmt.Println(dbConnection)
	db := config.ConnectToDB(dbConnection)
	defer db.Close()

	var option = -1
	for option != 0 {

		fmt.Println("~~~SELAMAT DATANG~~~")
		fmt.Println("Pilihan menu:")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Login")
		fmt.Println("3. Profil User")
		fmt.Println("4. Update Data User")
		fmt.Println("5. Hapus Akun User")
		fmt.Println("6. Top Up")
		fmt.Println("7. Transfer")
		fmt.Println("8. History Top Up")
		fmt.Println("9. History Transfer")
		fmt.Println("10. Cari User")
		fmt.Println("0. Log Out")

		fmt.Print("Masukkan Nomor Menu:")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Menu Sign Up")

			newUser := entities.User{}

			fmt.Print("Nama \t\t: ")
			fmt.Scanln(&newUser.NAMA)
			fmt.Print("No. Telp \t: ")
			fmt.Scanln(&newUser.NO_TELP)
			fmt.Print("Password \t: ")
			fmt.Scanln(&newUser.PASSWORD)

			var query = "insert into user (nama, no_telp, password) values (?, ?, ?)"
			statement, errPrepare := db.Prepare(query)

			if errPrepare != nil {
				fmt.Println("Error prepare insert", errPrepare.Error())
			}

			result, errExec := statement.Exec(newUser.NAMA, newUser.NO_TELP, newUser.PASSWORD)
			if errExec != nil {
				fmt.Println("Error exec insert", errExec.Error())
			} else {
				row, errRow := result.RowsAffected()
				if errRow != nil {
					fmt.Println("Error row insert", errRow.Error())
				}
				if row > 0 {
					fmt.Println("Success Insert Data")
				} else {
					fmt.Println("Gagal insert")
				}
			}

		case 2:
			fmt.Println("Menu Login")

		case 3:
			fmt.Println("Menu Profil User")

		case 4:
			fmt.Println("Menu Update Data")

		case 5:
			fmt.Println("Menu Hapus Akun")

		case 6:
			fmt.Println("Menu Top Up")

		case 7:
			fmt.Println("Menu Transfer")

		case 8:
			fmt.Println("Menu History Top Up")

		case 9:
			fmt.Println("Menu History Transfer")

		case 10:
			fmt.Println("Menu Cari User")

		case 0:
			fmt.Println("\nTerima Kasih Telah Bertransaksi")

		default:
			fmt.Println("PERINGATAN: NOMOR MENU TIDAK TERSEDIA!!!")

		}
	}

}

// row_user, row_balance, errInputUser := user.InputDataUser(db, newUser)
