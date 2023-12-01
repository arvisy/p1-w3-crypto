package main

import (
	"cryptoss/config"
	"cryptoss/entity"
	"cryptoss/handler"
	"fmt"
)

func main() {
	config.ConnectDB()

	for {
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var username, password string
			fmt.Println("Registrasi")
			fmt.Print("Masukkan username: \n")
			fmt.Scan(&username)
			fmt.Print("Masukkan password: \n")
			fmt.Scan(&password)

			// define entity
			user := entity.User{
				Username: username,
				Password: password,
			}

			// define check handler
			err := handler.Register(user)
			if err != nil {
				fmt.Println("Kesalahan saat registrasi", err)
			} else {
				fmt.Println("Berhasil registrasi")
			}
		case 2:
			var username, password string
			fmt.Println("Login")
			fmt.Print("Masukkan username: \n")
			fmt.Scan(&username)
			fmt.Print("Masukkan password: \n")
			fmt.Scan(&password)

			// define user auth dengan handler
			user, isAuthenticated := handler.Login(username, password)
			if isAuthenticated {
				fmt.Println("Login berhasil! Selamat datang", user.Username)
			} else {
				fmt.Println("Login gagal")
			}
		case 3:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid")

			// define user auth dengan handler
		}
	}
}
