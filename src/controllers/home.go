package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
	"strings"
)

//Carrega página inicial
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)
	primeiroNome := strings.Split(cookies.Nome, " ")

	utils.ExecutarTemplate(w, "home.html", struct {
		Cookies      modelos.PageCookies
		Pagina       string
		PrimeiroNome string
	}{
		PrimeiroNome: primeiroNome[0],
		Cookies:      cookies,
		Pagina:       "Home",
	})
}

//Carrega página inicial
func Handler404(w http.ResponseWriter, r *http.Request) {
	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "404.html", struct {
		Cookies modelos.PageCookies
		Pagina  string
	}{

		Cookies: cookies,
		Pagina:  "H4",
	})
}
