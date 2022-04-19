package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
)

//Carrega página de usuario
func CarregarPaginaDeConsultaDeCatraca(w http.ResponseWriter, r *http.Request) {
	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "consultar-catraca.html", struct {
		Cookies modelos.PageCookies
		Pagina  string
	}{
		Cookies: cookies,
		Pagina:  "Consultar Catraca",
	})
}
