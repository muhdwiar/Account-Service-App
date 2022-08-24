package main

import (
	"be11/project1/config"
	"be11/project1/controllers/user"
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

	db := config.ConnectToDB(dbConnection)
	defer db.Close()

	var User_Login = entities.User{}
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

		case 2:
			fmt.Println("Menu Login")

			fmt.Print("Silahkan Masukkan No. Telepon \t:")
			fmt.Scanln(&User_Login.NO_TELP)
			fmt.Print("Silahkan Masukkan Password \t:")
			fmt.Scanln(&User_Login.PASSWORD)

			temp_loginData, err_massage := user.LoginUser(db, User_Login)

			if err_massage != nil {
				fmt.Println("ERROR LOGIN :", err_massage.Error())
			} else {
				if (temp_loginData != entities.User{}) {
					User_Login = temp_loginData
					fmt.Println(User_Login)
					fmt.Print("BERHASIL LOGIN\n\n")
				} else {
					User_Login = entities.User{}
					fmt.Print("AKUN TIDAK ADA SILHAKAN MELAKUKAN SIGN UP\n\n")
				}
			}

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
