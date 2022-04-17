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
	"strings"

	"github.com/gorilla/mux"
)

//Carrega página principal
func CarregarPaginaInicialBase(w http.ResponseWriter, r *http.Request) {
	//recupera url query strings
	strCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	strCategoriaNoSpace := strings.ReplaceAll(strCategoria, " ", "+")
	strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")

	// define urls das api
	url := fmt.Sprintf("%s/base?categoria=%s&cliente=%s", config.APIURL, strCategoriaNoSpace, strClienteNoSpace)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlCategorias := fmt.Sprintf("%s/categorias", config.APIURL)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var posts []modelos.Post
	if erro = json.NewDecoder(response.Body).Decode(&posts); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos clientes
	responseClientes, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientes, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientes.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientes)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(responseClientes.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api das categorias
	responseCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCategorias)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(responseCategorias.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos Sites
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

	utils.ExecutarTemplate(w, "base.html", struct {
		Posts     []modelos.Post
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Cookies   modelos.PageCookies
		Pagina    string
	}{
		Posts:     posts,
		Cliente:   cliente,
		Categoria: categoria,
		Cookies:   cookies,
		Pagina:    "Página GSA",
		//Pagina:    "Chamado: " + chamado.Chamado,
	})
}

//Carrega página principal
func CarregarPaginaBuscaCatOuCliente(w http.ResponseWriter, r *http.Request) {
	//recupera url query strings
	strCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	strCategoriaNoSpace := strings.ReplaceAll(strCategoria, " ", "+")
	strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")

	// define urls das api
	url := fmt.Sprintf("%s/base/busca?categoria=%s&cliente=%s", config.APIURL, strCategoriaNoSpace, strClienteNoSpace)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlCategorias := fmt.Sprintf("%s/categorias", config.APIURL)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	//requisição para a api da base
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var posts []modelos.Post
	if erro = json.NewDecoder(response.Body).Decode(&posts); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos clientes
	responseClientes, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientes, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientes.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientes)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(responseClientes.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api das categorias
	responseCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCategorias)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(responseCategorias.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos Sites
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

	utils.ExecutarTemplate(w, "base_busca.html", struct {
		Posts     []modelos.Post
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Cookies   modelos.PageCookies
		Pagina    string
	}{
		Posts:     posts,
		Cliente:   cliente,
		Categoria: categoria,
		Cookies:   cookies,
		Pagina:    "Página GSA",
		//Pagina:    "Chamado: " + chamado.Chamado,
	})
}

//Carrega tela do formulario de criação de publicação.
func CarregarTelaDeCriarPublicacao(w http.ResponseWriter, r *http.Request) {

	// define urls das api
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlCategorias := fmt.Sprintf("%s/categorias", config.APIURL)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	//requisição para a api dos clientes
	responseClientes, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientes, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientes.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientes)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(responseClientes.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api das categorias
	responseCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCategorias)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(responseCategorias.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos Sites
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

	utils.ExecutarTemplate(w, "criar-publicacao.html", struct {
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Site      []modelos.Site
		Cookies   modelos.PageCookies
		Pagina    string
	}{

		Cliente:   cliente,
		Categoria: categoria,
		Site:      site,
		Cookies:   cookies,
		Pagina:    "Mapa Operações",
	})
}

//Carrega tela do formulario de criação de publicação.
func CarregarTelaDeCriarCategoria(w http.ResponseWriter, r *http.Request) {

	// define urls das api
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlCategorias := fmt.Sprintf("%s/categorias", config.APIURL)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	//requisição para a api dos clientes
	responseClientes, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientes, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientes.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientes)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(responseClientes.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api das categorias
	responseCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCategorias)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(responseCategorias.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos Sites
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

	utils.ExecutarTemplate(w, "criar-categoria.html", struct {
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Site      []modelos.Site
		Cookies   modelos.PageCookies
		Pagina    string
	}{

		Cliente:   cliente,
		Categoria: categoria,
		Site:      site,
		Cookies:   cookies,
		Pagina:    "Mapa Operações",
	})
}

//Carrega pagina de edição
func CarregarPaginaDePublicacao(w http.ResponseWriter, r *http.Request) {
	//recebe id através da url
	parametros := mux.Vars(r)
	postID, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// define urls da api
	url := fmt.Sprintf("%s/base/%d", config.APIURL, postID)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlCategorias := fmt.Sprintf("%s/categorias", config.APIURL)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	//requisição para a api da base de conhecimento
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	var post modelos.Post
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if erro = json.NewDecoder(response.Body).Decode(&post); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos clientes
	responseClientes, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientes, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientes.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientes)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(responseClientes.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api das categorias
	responseCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCategorias)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(responseCategorias.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos Sites
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

	utils.ExecutarTemplate(w, "post.html", struct {
		Post      modelos.Post
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Cookies   modelos.PageCookies
		Pagina    string
	}{
		Post:      post,
		Cliente:   cliente,
		Categoria: categoria,
		Cookies:   cookies,
		Pagina:    "Página GSA",
	})
}

//Carrega pagina de edição
func CarregarPaginaDeEdicaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	//recebe id através da url
	parametros := mux.Vars(r)
	postID, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// define urls da api
	url := fmt.Sprintf("%s/base/%d", config.APIURL, postID)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlCategorias := fmt.Sprintf("%s/categorias", config.APIURL)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	var post modelos.Post
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if erro = json.NewDecoder(response.Body).Decode(&post); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos clientes
	responseClientes, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientes, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientes.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientes)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(responseClientes.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api das categorias
	responseCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCategorias)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(responseCategorias.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos Sites
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

	utils.ExecutarTemplate(w, "editar-publicacao.html", struct {
		Post      modelos.Post
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Site      []modelos.Site
		Cookies   modelos.PageCookies
		Pagina    string
	}{
		Post:      post,
		Cliente:   cliente,
		Categoria: categoria,
		Site:      site,
		Cookies:   cookies,
		Pagina:    "Página GSA",
	})
}

//Carrega pagina de edição
func CarregarPaginaDeEdicaoDeCategorias(w http.ResponseWriter, r *http.Request) {
	//recebe id através da url
	parametros := mux.Vars(r)
	catID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// define urls da api
	url := fmt.Sprintf("%s/categorias/%d", config.APIURL, catID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	var categoria modelos.Post_Categoria
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if erro = json.NewDecoder(response.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "editar-categoria.html", struct {
		Categoria modelos.Post_Categoria
		Cookies   modelos.PageCookies
		Pagina    string
	}{
		Categoria: categoria,
		Cookies:   cookies,
		Pagina:    "Página GSA",
	})
}

//Carrega pagina de edição
func CarregarPaginaDeEdicaoDeClientes(w http.ResponseWriter, r *http.Request) {
	//recebe id através da url
	parametros := mux.Vars(r)
	catID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// define urls da api
	url := fmt.Sprintf("%s/clientes/%d", config.APIURL, catID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	var cliente modelos.Cliente
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if erro = json.NewDecoder(response.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "editar-cliente.html", struct {
		Cliente modelos.Cliente
		Cookies modelos.PageCookies
		Pagina  string
	}{
		Cliente: cliente,
		Cookies: cookies,
		Pagina:  "Página GSA",
	})
}
