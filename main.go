package main

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/router"
	"app/src/utils"
	"fmt"
	"log"
	"net/http"
)

//Gerar SecretKey
/*
func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)

	BlockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(BlockKey)
}

*/
func main() {

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Rodando o app na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
