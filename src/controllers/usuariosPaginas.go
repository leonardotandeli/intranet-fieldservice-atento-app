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
	"strconv"

	"github.com/gorilla/mux"
)

//Carrega pagina de alteração de senha
func CarregarPaginaDeEdicaoDeSenha(w http.ResponseWriter, r *http.Request) {

	//recebe id através da url
	parametros := mux.Vars(r)
	usuariocID, erro := strconv.ParseUint(parametros["usuariocId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// define url da api
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuariocID)

	//faz a requisição para a api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	var usuario modelos.Usuario
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if usuariocID == cookies.UsuarioID {

		utils.ExecutarTemplate(w, "alterarSenhaUsuario.html", struct {
			Usuario modelos.Usuario
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Usuario: usuario,
			Cookies: cookies,
			Pagina:  "Alterar Senha",
		})
	} else if cookies.V_USUARIOS == "S" {
		utils.ExecutarTemplate(w, "alterarSenhaUsuario.html", struct {
			Usuario modelos.Usuario
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Usuario: usuario,
			Cookies: cookies,
			Pagina:  "Alterar Senha",
		})
	} else {

		utils.ExecutarTemplate(w, "acessoNegado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Página Inicial",
		})
	}

}

//Carrega página de usuários
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	// define url da api
	url := fmt.Sprintf("%s/usuarios", config.APIURL)

	//faz a requisição para a api
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var usuarios []modelos.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_USUARIOS == "S" {
		utils.ExecutarTemplate(w, "usuarios.html", struct {
			Usuarios []modelos.Usuario
			Cookies  modelos.PageCookies
			SiteNome string
			Pagina   string
		}{
			Usuarios: usuarios,
			Cookies:  cookies,
			Pagina:   "Usuários",
		})

	} else {

		utils.ExecutarTemplate(w, "acessoNegado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Página Inicial",
		})

	}

}

//Carrega pagina de edição
func CarregarPaginaDeEdicaoDeUsuarios(w http.ResponseWriter, r *http.Request) {

	//recebe id através da url
	parametros := mux.Vars(r)
	usuariocID, erro := strconv.ParseUint(parametros["usuariocId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	// define urls da api
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuariocID)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	//faz a requisição para a api usuarios
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	var usuario modelos.Usuario
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//faz a requisição para a api sites
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

	//condicional de acesso a página
	if cookies.V_USUARIOS == "S" {

		utils.ExecutarTemplate(w, "alterarUsuario.html", struct {
			Usuario modelos.Usuario
			Site    []modelos.Site
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Usuario: usuario,
			Site:    site,
			Cookies: cookies,
			Pagina:  "Alterar Usuário",
			//Pagina:    "Chamado: " + chamado.Chamado,
		})

	} else {
		utils.ExecutarTemplate(w, "acessoNegado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Página Inicial",
		})

	}

}

//Carrega página de usuario
func CarregarPaginaDeCadastroDeUsuarios(w http.ResponseWriter, r *http.Request) {

	// define urls da api
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	//faz a requisição para a api sites
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

	//condicional de acesso a página
	if cookies.V_USUARIOS == "S" {

		utils.ExecutarTemplate(w, "adicionarUsuario.html", struct {
			Site    []modelos.Site
			Cookies modelos.PageCookies
			Pagina  string
		}{

			Site:    site,
			Cookies: cookies,
			Pagina:  "Criar novo usuário",
		})

	} else {
		utils.ExecutarTemplate(w, "acessoNegado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Página Inicial",
		})

	}
}

//Carrega página de usuario
func CarregarPaginaDeCadastroDeUsuariosExcel(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_USUARIOS == "S" {

		utils.ExecutarTemplate(w, "adicionarUsuarioEmMassa.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{

			Cookies: cookies,
			Pagina:  "Criar novo usuário",
		})

	} else {
		utils.ExecutarTemplate(w, "acessoNegado.html", struct {
			Cookies modelos.PageCookies
			Pagina  string
		}{
			Cookies: cookies,
			Pagina:  "Página Inicial",
		})

	}
}
