package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
)

//Carrega página do mapa de operações
func CarregarPaginaDeSalasDeCrise(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	if cookies.V_MAPA_OPERACIONAL == "S" {

		utils.ExecutarTemplate(w, "salas-crise.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{

			Cookies: cookies,
			Pagina:  "Salas de Crise",
		})
	} else {

		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Acesso Negado",
		})
	}
}
