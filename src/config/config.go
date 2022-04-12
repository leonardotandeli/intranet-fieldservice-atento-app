package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// representa url da api
	APIURL = ""
	//porta da api
	Porta = 0
	//utilizado para autenticar o cookie
	HashKey []byte
	// cripitografa os dados do cookie
	BlockKey []byte
)

// Inicializa as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}
	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
