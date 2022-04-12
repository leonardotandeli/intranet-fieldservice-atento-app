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
	strCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	strCliente := strings.ToLower(r.URL.Query().Get("cliente"))

	strCategoriaNoSpace := strings.ReplaceAll(strCategoria, " ", "+")
	strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")

	url := fmt.Sprintf("%s/base?categoria=%s&cliente=%s", config.APIURL, strCategoriaNoSpace, strClienteNoSpace)
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
	url2 := fmt.Sprintf("%s/clientes", config.APIURL)
	response2, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url2, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response2.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(response2.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url3 := fmt.Sprintf("%s/categorias", config.APIURL)
	response3, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url3, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response3.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(response3.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]

	utils.ExecutarTemplate(w, "base.html", struct {
		Posts []modelos.Post

		Cliente            []modelos.Cliente
		Categoria          []modelos.Post_Categoria
		UsuarioID          uint64
		Nome               string
		Login_NT           string
		RE                 string
		Cargo              string
		V_USUARIOS         string
		V_BDC_POSTS        string
		V_BDC_ADM          string
		V_IMDB             string
		V_GSA              string
		V_MAPA_OPERACIONAL string
		SiteNome           string
		Pagina             string
	}{
		Posts: posts,

		Cliente:            cliente,
		Categoria:          categoria,
		UsuarioID:          usuarioID,
		Nome:               nomeUsuario,
		Login_NT:           login_nt,
		RE:                 re,
		Cargo:              cargo,
		V_USUARIOS:         v_usuarios,
		V_BDC_POSTS:        v_bdc_posts,
		V_BDC_ADM:          v_bdc_adm,
		V_IMDB:             v_imdb,
		V_GSA:              v_gsa,
		V_MAPA_OPERACIONAL: v_mapa_operacional,
		SiteNome:           site_nome,
		Pagina:             "Página GSA",
		//Pagina:    "Chamado: " + chamado.Chamado,
	})
}

//Carrega página principal
func CarregarPaginaBuscaCatOuCliente(w http.ResponseWriter, r *http.Request) {
	strCategoria := strings.ToLower(r.URL.Query().Get("categoria"))
	strCliente := strings.ToLower(r.URL.Query().Get("cliente"))

	strCategoriaNoSpace := strings.ReplaceAll(strCategoria, " ", "+")
	strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")

	url := fmt.Sprintf("%s/base/busca?categoria=%s&cliente=%s", config.APIURL, strCategoriaNoSpace, strClienteNoSpace)
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
	url2 := fmt.Sprintf("%s/clientes", config.APIURL)
	response2, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url2, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response2.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(response2.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url3 := fmt.Sprintf("%s/categorias", config.APIURL)
	response3, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url3, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response3.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(response3.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]

	utils.ExecutarTemplate(w, "base_busca.html", struct {
		Posts              []modelos.Post
		Cliente            []modelos.Cliente
		Categoria          []modelos.Post_Categoria
		UsuarioID          uint64
		Nome               string
		Login_NT           string
		RE                 string
		Cargo              string
		V_USUARIOS         string
		V_BDC_POSTS        string
		V_BDC_ADM          string
		V_IMDB             string
		V_GSA              string
		V_MAPA_OPERACIONAL string
		SiteNome           string
		Pagina             string
	}{
		Posts:              posts,
		Cliente:            cliente,
		Categoria:          categoria,
		UsuarioID:          usuarioID,
		Nome:               nomeUsuario,
		Login_NT:           login_nt,
		RE:                 re,
		Cargo:              cargo,
		V_USUARIOS:         v_usuarios,
		V_BDC_POSTS:        v_bdc_posts,
		V_BDC_ADM:          v_bdc_adm,
		V_IMDB:             v_imdb,
		V_GSA:              v_gsa,
		V_MAPA_OPERACIONAL: v_mapa_operacional,
		SiteNome:           site_nome,
		Pagina:             "Página GSA",
		//Pagina:    "Chamado: " + chamado.Chamado,
	})
}

//Carrega pagina de edição
func CarregarPaginaDePublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	postID, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(postID)
	url := fmt.Sprintf("%s/base/%d", config.APIURL, postID)
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

	url2 := fmt.Sprintf("%s/clientes", config.APIURL)
	response2, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url2, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response2.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(response2.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url3 := fmt.Sprintf("%s/categorias", config.APIURL)
	response3, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url3, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response3.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(response3.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]
	utils.ExecutarTemplate(w, "post.html", struct {
		Post               modelos.Post
		Cliente            []modelos.Cliente
		Categoria          []modelos.Post_Categoria
		UsuarioID          uint64
		Nome               string
		Login_NT           string
		RE                 string
		Cargo              string
		V_USUARIOS         string
		V_BDC_POSTS        string
		V_BDC_ADM          string
		V_IMDB             string
		V_GSA              string
		V_MAPA_OPERACIONAL string
		SiteNome           string
		Pagina             string
	}{
		Post:               post,
		Cliente:            cliente,
		Categoria:          categoria,
		UsuarioID:          usuarioID,
		Nome:               nomeUsuario,
		Login_NT:           login_nt,
		RE:                 re,
		Cargo:              cargo,
		V_USUARIOS:         v_usuarios,
		V_BDC_POSTS:        v_bdc_posts,
		V_BDC_ADM:          v_bdc_adm,
		V_IMDB:             v_imdb,
		V_GSA:              v_gsa,
		V_MAPA_OPERACIONAL: v_mapa_operacional,
		SiteNome:           site_nome,
		Pagina:             "Página GSA",
		//Pagina:    "Chamado: " + chamado.Chamado,
	})
}

//Carrega pagina de edição
func CarregarPaginaDeEdicaoDePublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	postID, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(postID)
	url := fmt.Sprintf("%s/base/%d", config.APIURL, postID)
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

	url2 := fmt.Sprintf("%s/clientes", config.APIURL)
	response2, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url2, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response2.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(response2.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url3 := fmt.Sprintf("%s/categorias", config.APIURL)
	response3, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url3, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response3.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var categoria []modelos.Post_Categoria
	if erro = json.NewDecoder(response3.Body).Decode(&categoria); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url4 := fmt.Sprintf("%s/sites", config.APIURL)
	response4, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url4, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response4.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var site []modelos.Site
	if erro = json.NewDecoder(response4.Body).Decode(&site); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]
	utils.ExecutarTemplate(w, "editar-publicacao.html", struct {
		Post               modelos.Post
		Cliente            []modelos.Cliente
		Categoria          []modelos.Post_Categoria
		Site               []modelos.Site
		UsuarioID          uint64
		Nome               string
		Login_NT           string
		RE                 string
		Cargo              string
		V_USUARIOS         string
		V_BDC_POSTS        string
		V_BDC_ADM          string
		V_IMDB             string
		V_GSA              string
		V_MAPA_OPERACIONAL string
		SiteNome           string
		Pagina             string
	}{
		Post:               post,
		Cliente:            cliente,
		Categoria:          categoria,
		Site:               site,
		UsuarioID:          usuarioID,
		Nome:               nomeUsuario,
		Login_NT:           login_nt,
		RE:                 re,
		Cargo:              cargo,
		V_USUARIOS:         v_usuarios,
		V_BDC_POSTS:        v_bdc_posts,
		V_BDC_ADM:          v_bdc_adm,
		V_IMDB:             v_imdb,
		V_GSA:              v_gsa,
		V_MAPA_OPERACIONAL: v_mapa_operacional,
		SiteNome:           site_nome,
		Pagina:             "Página GSA",
		//Pagina:    "Chamado: " + chamado.Chamado,
	})
}

/*
//Carrega tela do formulario
func CarregarTelaDeCriarPublicacao(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	busca := strings.ToLower(r.URL.Query().Get("a"))
	aSite, _ := cookie["asite"]
	utils.ExecutarTemplate(w, "criar-publicacao.html", struct {
		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Pagina    string
		Busca     string
		ASite     string
	}{
		UsuarioID: usuarioID,
		Nome:      nomeUsuario,
		Perfil:    perfilUsuario,
		Pagina:    "Formulário de Troca de Equipamentos",
		Busca:     busca,
		ASite:     aSite,
	})
}



// CarregarPaginaDeUsuarios carrega a página com os usuários que atendem o filtro passado
func CarregarPaginaDeBusca(w http.ResponseWriter, r *http.Request) {
	nomeDoc := strings.ToLower(r.URL.Query().Get("a"))
	url := fmt.Sprintf("%s/buscar?a=%s", config.APIURL, nomeDoc)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	cookie, _ := cookies.Ler(r)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	aSite, _ := cookie["asite"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	busca := strings.ToLower(r.URL.Query().Get("a"))

	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "busca.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
		LoginNT     string
		Perfil      string
		ASite       string
		Busca       string
		Nome        string
		Pagina      string
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
		Busca:       busca,
		Nome:        nomeUsuario,
		Perfil:      perfilUsuario,
		ASite:       aSite,
		//Pagina:    "Chamado: " + chamado.Chamado,
	})

}
*/
