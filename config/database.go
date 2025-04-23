package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dns := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal("Erro ao abrir conexão no banco de dados: ", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
