package main

import (
	"be11/project1/config"
	"be11/project1/controllers/transaksi"
	"be11/project1/controllers/user"
	"be11/project1/entities"
	"fmt"
	"os"
	"strings"

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

			if (User_Login == entities.User{}) {
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

			} else {
				fmt.Print("\nANDA SUDAH LOGIN, SILAHKAN LOG OUT DAHULU\n\n")
			}

		case 2:
			if (User_Login == entities.User{}) {

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

			} else {
				fmt.Print("\nANDA SUDAH LOGIN, SILAHKAN LOG OUT DAHULU\n\n")
			}

		case 3:
			if (User_Login != entities.User{}) {
				fmt.Println("Profil User")
				// fmt.Println(User_Login)
				temp_prof, err_profile := user.ReadProfile(db, User_Login)
				if err_profile != nil {
					fmt.Println("Gagal tampilkan data", err_profile.Error())
				} else {
					fmt.Println("ID\t\t: ", temp_prof.ID, "\nNama\t\t: ", temp_prof.NAMA, "\nNo.Telp\t\t: ", temp_prof.NO_TELP,
						"\nMember sejak\t: ", temp_prof.CREATED_AT.Format("2006-01-02 15:04:05"),
						"\nUpdate\t\t: ", temp_prof.UPDATED_AT.Format("2006-01-02 15:04:05"), "\nSaldo\t\t: ", temp_prof.BALANCE.SALDO)
				}
			} else {
				fmt.Print("\nANDA BELUM LOGIN\n\n")
			}

		case 4:
			if (User_Login != entities.User{}) {
				fmt.Println("Menu Update Data")
				var User = entities.User{}

				fmt.Print("Nama \t\t: ")
				fmt.Scanln(&User.NAMA)
				fmt.Print("No. Telp \t: ")
				fmt.Scanln(&User.NO_TELP)
				fmt.Print("Password \t: ")
				fmt.Scanln(&User.PASSWORD)

				temp_updateprof, err_updateproff := user.UpdateProfile(db, User_Login, User)
				if err_updateproff != nil {
					fmt.Println("Gagal masukan data", err_updateproff.Error())
				} else {
					if temp_updateprof > 0 {
						fmt.Println("Success Insert Data")
					} else {
						fmt.Println("Gagal insert")
					}
				}

			} else {
				fmt.Print("\nANDA BELUM LOGIN\n\n")
			}

		case 5:
			if (User_Login != entities.User{}) {
				fmt.Println("Menu Hapus Akun")

				var choice string
				fmt.Print("APAKAH ANDA YAKIN INGIN MENGHAPUS AKUN ANDA (masukan Y untuk menghapus) : ")
				fmt.Scanln(&choice)
				if strings.ToUpper(choice) == "Y" {
					row_user, err := user.DeleteUser(db, User_Login)

					if err != nil {
						fmt.Println("ERROR DELETE USER:", err.Error())
					} else {
						if row_user > 0 {
							User_Login = entities.User{}
							fmt.Println("AKUN ANDA TELAH DIHAPUS, SILAHKAN LOGIN LAGI")
						} else {
							fmt.Println("GAGAL HAPUS AKUN")
						}
					}
				}

			} else {
				fmt.Print("\nANDA BELUM LOGIN\n\n")
			}

		case 6:
			if (User_Login != entities.User{}) {
				fmt.Println("Menu Top Up")
				var Transaksi = entities.Transaksi{}
				fmt.Print("Nominal \t: ")
				fmt.Scanln(&Transaksi.NOMINAL)

				temp, temp2, err := transaksi.TopUp(db, Transaksi, User_Login)
				if err != nil {
					fmt.Println("Gagal masukan data", err.Error())

				} else {
					if temp > 0 && temp2 > 0 {
						fmt.Println("Success Insert Data")
					} else {
						fmt.Println("Gagal insert")
					}
				}

			} else {
				fmt.Print("\nANDA BELUM LOGIN\n\n")
			}

		case 7:
			if (User_Login != entities.User{}) {
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
					row_user, row_penerima, row_trans, err := transaksi.Transaksi(db, trans_userLogin, userPenerima)

					if err != nil {
						fmt.Println(row_user, row_penerima, row_trans)
						fmt.Println("ERROR TRANSFER :", err.Error(), "\n ")
					} else {
						if row_user > 0 && row_penerima > 0 && row_trans > 0 {
							fmt.Println("BERHASIL TRANSFER\n ")
						} else {
							fmt.Println("GAGAL TRANSFER\n ")
						}
					}

				}

			} else {
				fmt.Print("\nANDA BELUM LOGIN\n\n")
			}

		case 8:
			if (User_Login != entities.User{}) {
				fmt.Println("Menu History Top Up")
				result, err := transaksi.HistoryTP(db, User_Login)
				if err != nil {
					fmt.Println("error membaca data dari database", err)
				} else {
					for _, v := range result {
						fmt.Println("Nominal :", v.NOMINAL, "\tPada Tgl: ", v.CREATED_AT.Format("2006-01-02 15:04:05"))
					}
				}
			} else {
				fmt.Print("\nANDA BELUM LOGIN\n\n")
			}

		case 9:
			if (User_Login != entities.User{}) {
				fmt.Println("Menu History Transfer")

				var opsiTrans int
				fmt.Println("1. Data Sebagai Pengirim")
				fmt.Println("2. Data Sebagai Penerima")
				fmt.Print("Masukkan Opsi :")
				fmt.Scanln(&opsiTrans)

				switch opsiTrans {
				case 1:
					data_sender, err_sender := transaksi.GetTransfer_asSender(db, User_Login)

					if err_sender != nil {
						fmt.Println("ERROR READ HISTORY TRANSFER :", err_sender.Error())
					} else {
						for _, v := range data_sender {
							fmt.Println("Nama :", v.USER.NAMA, "\tNo.Telp :", v.USER.NO_TELP,
								"\tNominal :", v.NOMINAL, "\tMember Sejak :", v.CREATED_AT.Format("2006-01-02 15:04:05"))
						}
					}
				case 2:
					data_receiver, err_receiver := transaksi.GetTransfer_asReceiver(db, User_Login)

					if err_receiver != nil {
						fmt.Println("ERROR READ HISTORY TRANSFER :", err_receiver.Error())
					} else {
						for _, v := range data_receiver {
							fmt.Println("Nama :", v.USER.NAMA, "\tNo.Telp :", v.USER.NO_TELP,
								"\tNominal :", v.NOMINAL, "\tMember Sejak :", v.CREATED_AT.Format("2006-01-02 15:04:05"))
						}
					}
				}
				fmt.Println()
			} else {
				fmt.Print("\nANDA BELUM LOGIN\n\n")
			}

		case 10:
			fmt.Println("Menu Cari User")
			temp_user := entities.User{}
			fmt.Print("Masukkan No.Telp User :")
			fmt.Scanln(&temp_user.NO_TELP)

			data_user, err_user := user.FindUser(db, temp_user)

			if err_user != nil {
				fmt.Println("Gagal mencari user :", err_user.Error())
			} else {
				if (data_user == entities.User{}) {
					fmt.Println("USER TIDAK DITEMUKAN")

				} else {
					fmt.Println("\nNama\t\t: ", data_user.NAMA,
						"\nNo.Telp\t\t: ", data_user.NO_TELP, "\nMember sejak\t: ", data_user.CREATED_AT.Format("2006-01-02 15:04:05"), "\n ")
				}
			}

		case 0:
			fmt.Println("\nTerima Kasih Telah Bertransaksi")

		default:
			fmt.Println("PERINGATAN: NOMOR MENU TIDAK TERSEDIA!!!")

		}
	}

}
