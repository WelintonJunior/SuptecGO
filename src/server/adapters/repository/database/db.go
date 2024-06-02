package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	if err := godotenv.Load("../../../.env"); err != nil {
		panic(err)
	}

	userDB := os.Getenv("userDB")
	passDB := os.Getenv("passDB")
	databaseDB := os.Getenv("databaseDB")
	hostDB := os.Getenv("hostDB")
	portDB := os.Getenv("portDB")

	var conn string = fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable", userDB, passDB, databaseDB, hostDB, portDB)

	var err error
	DB, err = sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")
}
