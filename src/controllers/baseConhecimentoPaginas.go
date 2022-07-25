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

//CarregarPaginaInicialBase carrega a página inicial da base de conhecimento
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

	utils.ExecutarTemplate(w, "baseDeConhecimento.html", struct {
		PostsLimit []modelos.Post
		Cliente    []modelos.Cliente
		Categoria  []modelos.Post_Categoria
		Cookies    modelos.PageCookies
		Pagina     string
	}{
		PostsLimit: posts,
		Cliente:    cliente,
		Categoria:  categoria,
		Cookies:    cookies,
		Pagina:     "Base de Conhecimento",
	})
}

//CarregarPaginaPostsBase carrega a página inicial da base de conhecimento
func CarregarPaginaUltimosPosts(w http.ResponseWriter, r *http.Request) {

	// define urls das api
	url := fmt.Sprintf("%s/base", config.APIURL)
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

	utils.ExecutarTemplate(w, "baseDeConhecimentoPosts.html", struct {
		Posts     []modelos.Post
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Cookies   modelos.PageCookies
		Pagina    string
	}{
		Posts:     posts,
		Cliente:   cliente[1:],
		Categoria: categoria,
		Cookies:   cookies,
		Pagina:    "Base de Conhecimento",
	})
}

//CarregarPaginaBusca carrega página de busca
func CarregarPaginaBusca(w http.ResponseWriter, r *http.Request) {
	//recupera url query strings
	strBusca := strings.ToLower(r.URL.Query().Get("busca"))
	strBuscaNoSpace := strings.ReplaceAll(strBusca, " ", "+")

	// define urls das api
	url := fmt.Sprintf("%s/base/search?busca=%s", config.APIURL, strBuscaNoSpace)
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

	utils.ExecutarTemplate(w, "searchBaseDeConhecimento.html", struct {
		Posts []modelos.Post

		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Cookies   modelos.PageCookies
		Pagina    string
	}{
		Posts: posts,

		Cliente:   cliente[1:],
		Categoria: categoria,
		Cookies:   cookies,
		Pagina:    strBuscaNoSpace,
	})
}

//CarregarPaginaBuscaCatOuCliente carrega a pagina de busca por fltro de cliente ou categoria
func CarregarPaginaBuscaCliente(w http.ResponseWriter, r *http.Request) {
	//recupera url query strings
	// strCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	// strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	// strCategoriaNoSpace := strings.ReplaceAll(strCategoria, " ", "+")
	// strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")

	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// define urls das api
	url := fmt.Sprintf("%s/base/busca?cliente=%d", config.APIURL, clienteID)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlCategorias := fmt.Sprintf("%s/categorias/cliente/%d", config.APIURL, clienteID)
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

	var (
		urlClientesUnique string
	)
	if clienteID != 0 {
		urlClientesUnique = fmt.Sprintf("%s/clientes/%d", config.APIURL, clienteID)

	} else {
		urlClientesUnique = fmt.Sprintf("%s/clientes/2", config.APIURL)
	}

	//requisição para a api dos clientes
	responseClientesUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientesUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientesUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientesUnique)
		return
	}
	var clienteUnique modelos.Cliente
	if erro = json.NewDecoder(responseClientesUnique.Body).Decode(&clienteUnique); erro != nil {
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

	utils.ExecutarTemplate(w, "baseDeConhecimentoCliente.html", struct {
		Posts         []modelos.Post
		ClienteUnique modelos.Cliente
		Cliente       []modelos.Cliente
		Categoria     []modelos.Post_Categoria
		Cookies       modelos.PageCookies
		Pagina        string
	}{
		Posts:         posts,
		ClienteUnique: clienteUnique,
		Cliente:       cliente[1:],
		Categoria:     categoria,
		Cookies:       cookies,
		Pagina:        clienteUnique.NOME,
	})
}

//CarregarPaginaBuscaCatOuCliente carrega a pagina de busca por fltro de cliente ou categoria
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

	var (
		urlClientesUnique string
	)
	if strClienteNoSpace != "" {
		urlClientesUnique = fmt.Sprintf("%s/clientes/%s", config.APIURL, strClienteNoSpace)

	} else {
		urlClientesUnique = fmt.Sprintf("%s/clientes/2", config.APIURL)
	}

	//requisição para a api dos clientes
	responseClientesUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientesUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientesUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientesUnique)
		return
	}
	var clienteUnique modelos.Cliente
	if erro = json.NewDecoder(responseClientesUnique.Body).Decode(&clienteUnique); erro != nil {
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

	utils.ExecutarTemplate(w, "baseDeConhecimentoCliente.html", struct {
		Posts         []modelos.Post
		ClienteUnique modelos.Cliente
		Cliente       []modelos.Cliente
		Categoria     []modelos.Post_Categoria
		Cookies       modelos.PageCookies
		Pagina        string
	}{
		Posts:         posts,
		ClienteUnique: clienteUnique,
		Cliente:       cliente[1:],
		Categoria:     categoria,
		Cookies:       cookies,
		Pagina:        clienteUnique.NOME,
	})
}

//arregarPaginaBuscaCategoria carrega página de busca por categoria
func CarregarPaginaBuscaCategoria(w http.ResponseWriter, r *http.Request) {
	//recupera url query strings

	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	categoriaID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// strCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	// strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	// strCategoriaNoSpace := strings.ReplaceAll(strCategoria, " ", "+")
	// strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")

	// define urls das api
	url := fmt.Sprintf("%s/base/busca-cat-cliente?categoria=%d&cliente=%d", config.APIURL, categoriaID, clienteID)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlSubCategorias := fmt.Sprintf("%s/categorias/subcategoria/%d", config.APIURL, categoriaID)
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

	var (
		urlCatUnique string
	)

	urlCatUnique = fmt.Sprintf("%s/categorias/%d", config.APIURL, categoriaID)

	//requisição para a api dos clientes
	responseCatUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCatUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCatUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCatUnique)
		return
	}
	var catUnique modelos.Post_Categoria
	if erro = json.NewDecoder(responseCatUnique.Body).Decode(&catUnique); erro != nil {
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

	var (
		urlClientesUnique string
	)
	if clienteID != 0 {
		urlClientesUnique = fmt.Sprintf("%s/clientes/%d", config.APIURL, clienteID)

	} else {
		urlClientesUnique = fmt.Sprintf("%s/clientes/2", config.APIURL)
	}

	//requisição para a api dos clientes
	responseClientesUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientesUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientesUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientesUnique)
		return
	}
	var clienteUnique modelos.Cliente
	if erro = json.NewDecoder(responseClientesUnique.Body).Decode(&clienteUnique); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api das categorias
	responseCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlSubCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCategorias)
		return
	}
	var subcategoria []modelos.Post_SubCategoria
	if erro = json.NewDecoder(responseCategorias.Body).Decode(&subcategoria); erro != nil {
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

	utils.ExecutarTemplate(w, "baseDeConhecimentoCategoria.html", struct {
		Posts         []modelos.Post
		ClienteUnique modelos.Cliente
		CatUnique     modelos.Post_Categoria
		Cliente       []modelos.Cliente
		SubCategoria  []modelos.Post_SubCategoria
		CategoriaID   uint64
		Cookies       modelos.PageCookies
		Pagina        string
	}{
		Posts:         posts,
		ClienteUnique: clienteUnique,
		CatUnique:     catUnique,
		Cliente:       cliente[1:],
		SubCategoria:  subcategoria,
		CategoriaID:   categoriaID,
		Cookies:       cookies,
		Pagina:        catUnique.NOME,
	})
}

//arregarPaginaBuscaCategoria carrega página de busca por categoria
func CarregarPaginaBuscaSubCategoria(w http.ResponseWriter, r *http.Request) {
	//recupera url query strings

	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	categoriaID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	subCategoriaID, erro := strconv.ParseUint(parametros["subCatId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// strCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	// strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	// strCategoriaNoSpace := strings.ReplaceAll(strCategoria, " ", "+")
	// strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")

	// define urls das api
	url := fmt.Sprintf("%s/base/busca-subcat-cliente?subcategoria=%d&cliente=%d", config.APIURL, subCategoriaID, clienteID)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlSubCategorias := fmt.Sprintf("%s/categorias/subcategoria/%d", config.APIURL, categoriaID)
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

	var (
		urlCatUnique    string
		urlSubCatUnique string
	)

	urlCatUnique = fmt.Sprintf("%s/categorias/%d", config.APIURL, categoriaID)

	//requisição para a api dos clientes
	responseCatUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCatUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCatUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCatUnique)
		return
	}
	var catUnique modelos.Post_Categoria
	if erro = json.NewDecoder(responseCatUnique.Body).Decode(&catUnique); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	urlSubCatUnique = fmt.Sprintf("%s/subcategorias/%d", config.APIURL, subCategoriaID)
	//requisição para a api dos clientes
	responseSubCatUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlSubCatUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseSubCatUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseSubCatUnique)
		return
	}
	var subcatUnique modelos.Post_SubCategoria
	if erro = json.NewDecoder(responseSubCatUnique.Body).Decode(&subcatUnique); erro != nil {
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

	var (
		urlClientesUnique string
	)
	if clienteID != 0 {
		urlClientesUnique = fmt.Sprintf("%s/clientes/%d", config.APIURL, clienteID)

	} else {
		urlClientesUnique = fmt.Sprintf("%s/clientes/2", config.APIURL)
	}

	//requisição para a api dos clientes
	responseClientesUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientesUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientesUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientesUnique)
		return
	}
	var clienteUnique modelos.Cliente
	if erro = json.NewDecoder(responseClientesUnique.Body).Decode(&clienteUnique); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api das categorias
	responseCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlSubCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCategorias)
		return
	}
	var subcategoria []modelos.Post_SubCategoria
	if erro = json.NewDecoder(responseCategorias.Body).Decode(&subcategoria); erro != nil {
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

	utils.ExecutarTemplate(w, "baseDeConhecimentoSubCategoria.html", struct {
		Posts          []modelos.Post
		ClienteUnique  modelos.Cliente
		CatUnique      modelos.Post_Categoria
		SubCatUnique   modelos.Post_SubCategoria
		ClienteID      uint64
		CategoriaID    uint64
		SubCategoriaID uint64
		Cliente        []modelos.Cliente
		SubCategoria   []modelos.Post_SubCategoria
		Cookies        modelos.PageCookies
		Pagina         string
	}{
		Posts:          posts,
		ClienteUnique:  clienteUnique,
		CatUnique:      catUnique,
		SubCatUnique:   subcatUnique,
		ClienteID:      clienteID,
		CategoriaID:    categoriaID,
		SubCategoriaID: subCategoriaID,
		Cliente:        cliente[1:],
		SubCategoria:   subcategoria,
		Cookies:        cookies,
		Pagina:         catUnique.NOME,
	})
}

//CarregarTelaDeCriarPublicacao carrega a página da publicação
func CarregarPaginaDeCriarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	categoriaID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// define urls das api
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlCategorias := fmt.Sprintf("%s/categorias", config.APIURL)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)
	urlSubCategorias := fmt.Sprintf("%s/categorias/subcategoria/%d", config.APIURL, categoriaID)
	urlCatUnique := fmt.Sprintf("%s/categorias/%d", config.APIURL, categoriaID)
	urlClientesUnique := fmt.Sprintf("%s/clientes/%d", config.APIURL, clienteID)

	//requisição para a api das categorias
	responseSubCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlSubCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseSubCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseSubCategorias)
		return
	}
	var subcategoria []modelos.Post_SubCategoria
	if erro = json.NewDecoder(responseSubCategorias.Body).Decode(&subcategoria); erro != nil {
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

	//requisição para a api dos clientes
	responseClientesUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlClientesUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseClientesUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseClientesUnique)
		return
	}
	var clienteUnique modelos.Cliente
	if erro = json.NewDecoder(responseClientesUnique.Body).Decode(&clienteUnique); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos clientes
	responseCatUnique, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlCatUnique, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseCatUnique.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseCatUnique)
		return
	}
	var catUnique modelos.Post_Categoria
	if erro = json.NewDecoder(responseCatUnique.Body).Decode(&catUnique); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "adicionarPublicacao.html", struct {
		Cliente       []modelos.Cliente
		Categoria     []modelos.Post_Categoria
		Site          []modelos.Site
		SubCategoria  []modelos.Post_SubCategoria
		ClienteUnique modelos.Cliente
		CatUnique     modelos.Post_Categoria
		CategoriaID   uint64
		ClienteID     uint64
		Cookies       modelos.PageCookies
		Pagina        string
	}{

		Cliente:      cliente,
		Categoria:    categoria,
		Site:         site,
		SubCategoria: subcategoria,

		ClienteUnique: clienteUnique,
		CatUnique:     catUnique,
		CategoriaID:   categoriaID,
		ClienteID:     clienteID,
		Cookies:       cookies,
		Pagina:        "Criar Publicação",
	})
}

//CarregarTelaDeCriarCategoria carrega a página de formulario de criação de categoria.
func CarregarPaginaDeCriarCategoria(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
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

	utils.ExecutarTemplate(w, "adicionarCategoria.html", struct {
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Site      []modelos.Site
		ClienteID uint64
		Cookies   modelos.PageCookies
		Pagina    string
	}{

		Cliente:   cliente,
		Categoria: categoria,
		Site:      site,
		ClienteID: clienteID,
		Cookies:   cookies,
		Pagina:    "Criar nova categoria",
	})
}

//CarregarTelaDeCriarCategoria carrega a página de formulario de criação de categoria.
func CarregarPaginaDeCriarSubCategoria(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	clienteID, erro := strconv.ParseUint(parametros["clienteId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	catID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
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

	utils.ExecutarTemplate(w, "adicionarSubCategoria.html", struct {
		Cliente   []modelos.Cliente
		Categoria []modelos.Post_Categoria
		Site      []modelos.Site
		ClienteID uint64
		CatID     uint64
		Cookies   modelos.PageCookies
		Pagina    string
	}{

		Cliente:   cliente,
		Categoria: categoria,
		Site:      site,
		ClienteID: clienteID,
		CatID:     catID,
		Cookies:   cookies,
		Pagina:    "Criar nova categoria",
	})
}

//CarregarTelaDeCriarCliente carrega tela de criação de cliente.
func CarregarTelaDeCriarCliente(w http.ResponseWriter, r *http.Request) {

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

	utils.ExecutarTemplate(w, "adicionarCliente.html", struct {
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
		Pagina:    "Criar novo cliente",
	})
}

//CarregarPaginaDePublicacao carrega pagina da publicação
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
	urlSubCategorias := fmt.Sprintf("%s/subcategorias", config.APIURL)
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

	//requisição para a api das categorias
	responseSubCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlSubCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseSubCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseSubCategorias)
		return
	}
	var subcategoria []modelos.Post_SubCategoria
	if erro = json.NewDecoder(responseSubCategorias.Body).Decode(&subcategoria); erro != nil {
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

	utils.ExecutarTemplate(w, "publicacao.html", struct {
		Post         modelos.Post
		Cliente      []modelos.Cliente
		Categoria    []modelos.Post_Categoria
		SubCategoria []modelos.Post_SubCategoria
		Cookies      modelos.PageCookies
		Pagina       string
	}{
		Post:         post,
		Cliente:      cliente[1:],
		Categoria:    categoria,
		SubCategoria: subcategoria,
		Cookies:      cookies,
		Pagina:       post.TITULO,
	})
}

//CarregarPaginaDeEdicaoDePublicacao carrega pagina de alterar publicação
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
	urlSubCategorias := fmt.Sprintf("%s/subcategorias", config.APIURL)
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

	//requisição para a api das categorias
	responseSubCategorias, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlSubCategorias, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseSubCategorias.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseSubCategorias)
		return
	}
	var subcategoria []modelos.Post_SubCategoria
	if erro = json.NewDecoder(responseSubCategorias.Body).Decode(&subcategoria); erro != nil {
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

	utils.ExecutarTemplate(w, "alterarPublicacao.html", struct {
		Post         modelos.Post
		Cliente      []modelos.Cliente
		Categoria    []modelos.Post_Categoria
		SubCategoria []modelos.Post_SubCategoria
		Site         []modelos.Site
		Cookies      modelos.PageCookies
		Pagina       string
	}{
		Post:         post,
		Cliente:      cliente,
		Categoria:    categoria,
		SubCategoria: subcategoria,
		Site:         site,
		Cookies:      cookies,
		Pagina:       post.TITULO,
	})
}

//CarregarPaginaDeEdicaoDeCategorias carrega pagina de alterar categorias
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
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)

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

	utils.ExecutarTemplate(w, "alterarCategoria.html", struct {
		Categoria modelos.Post_Categoria
		Cookies   modelos.PageCookies
		Cliente   []modelos.Cliente
		Site      []modelos.Site
		Pagina    string
	}{
		Categoria: categoria,
		Cookies:   cookies,
		Cliente:   cliente,
		Site:      site,
		Pagina:    "Editar Categoria",
	})
}

//CarregarPaginaDeEdicaoDeClientes carrega pagina de alterar clientes
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

	utils.ExecutarTemplate(w, "alterarCliente.html", struct {
		Cliente modelos.Cliente
		Cookies modelos.PageCookies
		Pagina  string
	}{
		Cliente: cliente,
		Cookies: cookies,
		Pagina:  "Página GSA",
	})
}
