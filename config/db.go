package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DbUsername string
	DbPassword string
	DbHost     string
	DbName     string
	DbPort     string
}

// load env
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	return &Config{
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbHost:     dbHost,
		DbName:     dbName,
		DbPort:     dbPort,
	}

}

func MySQL(c *Config) (*sql.DB, error) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.DbUsername, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
	db, err := sql.Open("mysql", dbUri)

	if err != nil {
		return nil, err
	}

	fmt.Println("koneksi sukses")
	return db, nil
}
