package main

import (
	"be11/project1/config"
	"be11/project1/controllers/transaksi"
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
			newUser := entities.User{}

			fmt.Print("Nama \t\t: ")
			fmt.Scanln(&newUser.NAMA)
			fmt.Print("No. Telp \t: ")
			fmt.Scanln(&newUser.NO_TELP)
			fmt.Print("Password \t: ")
			fmt.Scanln(&newUser.PASSWORD)

			row_user, row_balance, err := user.Registrasi(db, newUser)

			if err != nil {
				fmt.Println("Gagal masukan data", err.Error())

			} else {
				if row_user > 0 && row_balance > 0 {
					fmt.Println("Success Insert Data")
				} else {
					fmt.Println("Gagal insert")
				}
			}

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
			trans_userLogin := entities.Transaksi{}
			userPenerima := entities.User{}

			fmt.Print("Masukkan No.Telp Penerima : ")
			fmt.Scanln(&userPenerima.NO_TELP)
			fmt.Print("Masukkan Nominal Transfer : ")
			fmt.Scanln(&trans_userLogin.NOMINAL)
			trans_userLogin.USER_ID = User_Login.ID

			if userPenerima.NO_TELP == User_Login.NO_TELP {
				fmt.Println("GAGAL TRANSFER, NO.TELP PENERIMA TIDAK BOLEH DIISI NO.TELP PENGIRIM\n ")
			} else {
				row_user, row_penerima, row_trans, err := transaksi.TransferBalance(db, trans_userLogin, userPenerima)

				if err != nil {
					fmt.Println("ERROR TRANSFER :", err.Error(), "\n ")
				} else {
					if row_user > 0 && row_penerima > 0 && row_trans > 0 {
						fmt.Println("BERHASIL TRANSFER\n ")
					} else {
						fmt.Println("GAGAL TRANSFER\n ")
					}
				}

			}

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
