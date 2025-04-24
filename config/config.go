package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Erro ao carregar o arquivo .env")
	} else {
		log.Println("Arquivo .env carregado com sucesso")
	}
}
