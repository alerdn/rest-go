package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	driverName := os.Getenv("DB_DRIVER")
	dns := os.Getenv("DATABASE_URL")

	DB, err = sql.Open(driverName, dns)
	if err != nil {
		log.Fatal("Erro ao abrir conexão no banco de dados: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco de dados: ", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
