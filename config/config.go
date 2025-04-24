package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type SMTPConfig struct {
	Host string
	Port int
	User string
	Pass string
}

var SMTP SMTPConfig

func LoadEnv() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Erro ao carregar o arquivo .env")
	} else {
		log.Println("Arquivo .env carregado com sucesso")
	}

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	SMTP = SMTPConfig{
		Host: os.Getenv("SMTP_HOST"),
		Port: port,
		User: os.Getenv("SMTP_USER"),
		Pass: os.Getenv("SMTP_PASS"),
	}
}
