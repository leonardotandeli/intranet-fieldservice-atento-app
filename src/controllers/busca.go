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

//Carrega página principal
func CarregarPaginaConsultaAD(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "consulta-ad-pagina.html", struct {
		Cookies modelos.PageCookies
		Pagina  string
	}{

		Cookies: cookies,
		Pagina:  "ConsultaAD",
	})
}

//Carrega página principal
func CarregarPaginaConsultaLAPS(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "consulta-laps-pagina.html", struct {
		Cookies modelos.PageCookies
		Pagina  string
	}{

		Cookies: cookies,
		Pagina:  "Consulta LAPS",
	})
}

//Carrega página principal
func ConsultaAD(w http.ResponseWriter, r *http.Request) {
	//parametros recebe dados através da url
	//parametros := mux.Vars(r)
	//loginNT := parametros["login"]

	loginNT := strings.ToLower(r.URL.Query().Get("login"))
	// define urls das api
	url := fmt.Sprintf("%s/ad/%s", config.APIURL, loginNT)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var dados modelos.DadosAD
	if erro = json.NewDecoder(response.Body).Decode(&dados); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	re := dados.LOGIN_NT
	reWithoutAB := strings.ReplaceAll(re, "AB", "")

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "consulta-ad.html", struct {
		DadosAD modelos.DadosAD
		RE      string
		Cookies modelos.PageCookies
		Pagina  string
	}{
		DadosAD: dados,
		RE:      reWithoutAB,
		Cookies: cookies,
		Pagina:  "DadosAD",
	})
}

//Carrega página principal
func ConsultaLAPS(w http.ResponseWriter, r *http.Request) {
	//parametros recebe dados através da url
	//parametros := mux.Vars(r)
	//locador := parametros["locador"]

	locador := strings.ToLower(r.URL.Query().Get("locador"))

	// define urls das api
	url := fmt.Sprintf("%s/laps/%s", config.APIURL, locador)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		http.Redirect(w, r, "/consulta/laps", http.StatusFound)
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var dados modelos.DadosLAPS
	if erro = json.NewDecoder(response.Body).Decode(&dados); erro != nil {
		http.Redirect(w, r, "/consulta/laps", http.StatusFound)
	}

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "consulta-laps.html", struct {
		DadosLAPS modelos.DadosLAPS
		RE        string
		Cookies   modelos.PageCookies
		Pagina    string
	}{
		DadosLAPS: dados,
		Cookies:   cookies,
		Pagina:    "DadosLaps",
	})
}
