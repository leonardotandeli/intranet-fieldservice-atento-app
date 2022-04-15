package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
)

//Carrega página inicial
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "home.html", struct {
		Cookies modelos.PageCookies
		Pagina  string
	}{

		Cookies: cookies,
		Pagina:  "Mapa Operações",
	})
}
