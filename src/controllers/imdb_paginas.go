package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
)

//Carrega página de usuario
func CarregarPaginaDeConsultaDeIMDB(w http.ResponseWriter, r *http.Request) {
	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_IMDB == "S" {

		utils.ExecutarTemplate(w, "consultar-imdb.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Consulta IMDB",
		})
	} else {

		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Página Inicial",
		})
	}
}
