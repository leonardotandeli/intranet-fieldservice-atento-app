package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
)

//Carrega página de layout
func CarregarPaginaMapaField(w http.ResponseWriter, r *http.Request) {
	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso
	if cookies.V_IMDB == "S" {

		utils.ExecutarTemplate(w, "mapa-field.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Mapeamento Field",
		})
	} else {

		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
			Pagina  string
			Cookies modelos.PageCookies
		}{
			Pagina:  "Página Inicial",
			Cookies: cookies,
		})
	}
}
