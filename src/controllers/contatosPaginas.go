package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
)

//CarregarPaginaDeContatosField página de contatos do Field
func CarregarPaginaDeContatosField(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_MAPA_OPERACIONAL == "S" {

		utils.ExecutarTemplate(w, "contatosField.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{

			Cookies: cookies,
			Pagina:  "Contatos",
		})
	} else {

		utils.ExecutarTemplate(w, "acessoNegado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Acesso Negado",
		})
	}
}

//CarregarPaginaDeContatosTsystems carrega página de contatos Tsystems
func CarregarPaginaDeContatosTsystems(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_MAPA_OPERACIONAL == "S" {

		utils.ExecutarTemplate(w, "contatosTsystems.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Contatos T-Systems",
		})
	} else {

		utils.ExecutarTemplate(w, "acessoNegado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Acesso Negado",
		})
	}
}
