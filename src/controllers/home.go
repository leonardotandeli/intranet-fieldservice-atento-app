package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/modelos"
	"app/src/requisicoes"
	"app/src/respostas"
	"app/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Carrega página inicial
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {

	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	//requisição para a api dos sites
	responseSites, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlSites, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseSites.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseSites)
		return
	}
	var site []modelos.Site
	if erro = json.NewDecoder(responseSites.Body).Decode(&site); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//função para inserir dados dos cookies armazenados durante o login

	cookies, _ := cookies.InserirDadosNaPagina(r)
	primeiroNome := strings.Split(cookies.Nome, " ")

	utils.ExecutarTemplate(w, "home.html", struct {
		Cookies      modelos.PageCookies
		Site         []modelos.Site
		Pagina       string
		PrimeiroNome string
	}{
		PrimeiroNome: primeiroNome[0],
		Cookies:      cookies,
		Site:         site[1:33],
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
