package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
)

//Carrega página do mapa de operações
func CarregarPaginaDeContatosField(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_MAPA_OPERACIONAL == "S" {

		utils.ExecutarTemplate(w, "contatos-field.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{

			Cookies: cookies,
			Pagina:  "Mapa Operações",
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

//Carrega página do mapa de operações
func CarregarPaginaDeContatosTsystems(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_MAPA_OPERACIONAL == "S" {

		utils.ExecutarTemplate(w, "contatos-tsystems.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Mapa Operações",
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
