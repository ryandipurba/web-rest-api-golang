package main

import (
	"book-catalog/config"
	"book-catalog/server"

	"github.com/go-playground/validator/v10"
)

func main() {
	validation := validator.New()
	cfg := config.LoadConfig()       //ngeload env
	dbInit, err := config.MySQL(cfg) //konek ke database dengan paramaeter yang ada di env
	if err != nil {
		panic(err)
	}

	server := server.NewServer(dbInit, validation) // memanggil koneksi

	server.ListenAndServer("8080")
}
