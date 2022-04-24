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

//Carrega página do mapa de operações
func CarregarPaginaMapa(w http.ResponseWriter, r *http.Request) {
	//recebe query strings da url
	strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	strSite := strings.ToLower(r.URL.Query().Get("site"))
	strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")
	strSiteNoSpace := strings.ReplaceAll(strSite, " ", "+")

	// define urls das api
	url := fmt.Sprintf("%s/mapa/operacoes?cliente=%s&site=%s", config.APIURL, strClienteNoSpace, strSiteNoSpace)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)

	//requisição para a api do mapa operacional
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var MapaOperacional []modelos.MapaOperacional
	if erro = json.NewDecoder(response.Body).Decode(&MapaOperacional); erro != nil {
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

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_MAPA_OPERACIONAL == "S" {
		utils.ExecutarTemplate(w, "mapa.html", struct {
			MapaOperacional []modelos.MapaOperacional
			Site            []modelos.Site
			Cliente         []modelos.Cliente
			Cookies         modelos.PageCookies
			Pagina          string
		}{
			MapaOperacional: MapaOperacional,
			Site:            site,
			Cliente:         cliente[1:],
			Cookies:         cookies,

			Pagina: "Mapa Operações",
		})
		//executa template da página de acesso negado.
	} else {
		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
			Pagina string
		}{
			Pagina: "Acesso Negado",
		})
	}
}

//Carrega página do mapa de operações
func CarregarPaginaMapaString(w http.ResponseWriter, r *http.Request) {
	//recebe query strings da url
	strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	strSite := strings.ToLower(r.URL.Query().Get("site"))
	strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")
	strSiteNoSpace := strings.ReplaceAll(strSite, " ", "+")

	// define urls das api
	url := fmt.Sprintf("%s/mapa/busca?cliente=%s&site=%s", config.APIURL, strClienteNoSpace, strSiteNoSpace)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)

	//requisição para a api do mapa operacional
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var MapaOperacional []modelos.MapaOperacional
	if erro = json.NewDecoder(response.Body).Decode(&MapaOperacional); erro != nil {
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

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	//condicional de acesso a página
	if cookies.V_MAPA_OPERACIONAL == "S" {
		utils.ExecutarTemplate(w, "mapa-string.html", struct {
			MapaOperacional []modelos.MapaOperacional
			Site            []modelos.Site
			Cliente         []modelos.Cliente
			Cookies         modelos.PageCookies
			Pagina          string
		}{
			MapaOperacional: MapaOperacional,
			Site:            site,
			Cliente:         cliente[1:],
			Cookies:         cookies,

			Pagina: "Busca Mapa Operações",
		})
		//executa template da página de acesso negado.
	} else {
		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
			Pagina string
		}{
			Pagina: "Acesso Negado",
		})
	}
}

//Carrega tela do formulario de adição de novos registros no mapa operacional
func CarregarTelaDoFormularioMapa(w http.ResponseWriter, r *http.Request) {

	// define urls das api
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlDacs := fmt.Sprintf("%s/dacs", config.APIURL)
	urlDominios := fmt.Sprintf("%s/dominios", config.APIURL)

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

	//requisição para a api dos Dacs
	responseDacs, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlDacs, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseDacs.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseDacs)
		return
	}
	var dac []modelos.Dac
	if erro = json.NewDecoder(responseDacs.Body).Decode(&dac); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos Domínios
	responseDominios, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlDominios, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseDominios.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseDominios)
		return
	}
	var dominio []modelos.Dominio
	if erro = json.NewDecoder(responseDominios.Body).Decode(&dominio); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "formulario-mapa.html", struct {
		Site    []modelos.Site
		Cliente []modelos.Cliente
		Dac     []modelos.Dac
		Dominio []modelos.Dominio
		Cookies modelos.PageCookies
		Pagina  string
	}{

		Site:    site,
		Cliente: cliente[1:],
		Dac:     dac,
		Dominio: dominio,
		Cookies: cookies,
		Pagina:  "Criar nova operação",
	})
}

//Carrega pagina de edição
func CarregarPaginaDeEdicaoMapa(w http.ResponseWriter, r *http.Request) {
	// recebe id da url
	parametros := mux.Vars(r)

	mapaID, erro := strconv.ParseUint(parametros["mapaId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// define urls das api
	url := fmt.Sprintf("%s/mapa/operacoes/%d", config.APIURL, mapaID)
	urlSites := fmt.Sprintf("%s/sites", config.APIURL)
	urlClientes := fmt.Sprintf("%s/clientes", config.APIURL)
	urlDacs := fmt.Sprintf("%s/dacs", config.APIURL)
	urlDominios := fmt.Sprintf("%s/dominios", config.APIURL)

	//requisição para a api do mapa operacional
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	var MapaOperacional modelos.MapaOperacional
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if erro = json.NewDecoder(response.Body).Decode(&MapaOperacional); erro != nil {
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

	//requisição para a api dos Dacs
	responseDacs, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlDacs, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseDacs.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseDacs)
		return
	}
	var dac []modelos.Dac
	if erro = json.NewDecoder(responseDacs.Body).Decode(&dac); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//requisição para a api dos Domínios
	responseDominios, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlDominios, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if responseDominios.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseDominios)
		return
	}
	var dominio []modelos.Dominio
	if erro = json.NewDecoder(responseDominios.Body).Decode(&dominio); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//função para inserir dados dos cookies armazenados durante o login
	cookies, _ := cookies.InserirDadosNaPagina(r)

	utils.ExecutarTemplate(w, "editar-mapa.html", struct {
		MapaOperacional modelos.MapaOperacional
		Site            []modelos.Site
		Cliente         []modelos.Cliente
		Dac             []modelos.Dac
		Dominio         []modelos.Dominio
		Cookies         modelos.PageCookies
		Pagina          string
	}{
		MapaOperacional: MapaOperacional,
		Site:            site,
		Cliente:         cliente[1:],
		Dac:             dac,
		Dominio:         dominio,
		Cookies:         cookies,
		Pagina:          "Edição de Operação",
	})
}
