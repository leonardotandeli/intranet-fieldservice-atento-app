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

//Carrega a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecutarTemplate(w, "login.html", nil)
}

//Carrega página de usuario
func CarregarPaginaDeCadastroDeUsuarios(w http.ResponseWriter, r *http.Request) {

	url2 := fmt.Sprintf("%s/sites", config.APIURL)
	response2, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url2, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response2.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var site []modelos.Site
	if erro = json.NewDecoder(response2.Body).Decode(&site); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]
	utils.ExecutarTemplate(w, "criar-usuario.html", struct {
		Usuarios                             []modelos.Usuario
		UsuarioID                            uint64
		Site                                 []modelos.Site
		Nome                                 string
		Login_NT                             string
		RE                                   string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{
		UsuarioID:                            usuarioID,
		Site:                                 site,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		RE:                                   re,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
		Pagina:                               "Mapa Operações",
	})
}

//Carrega página principal
func CarregarPaginaDeConfiguracoes(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios", config.APIURL)
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

	cookie, _ := cookies.Ler(r)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	aSite, _ := cookie["asite"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "configuracoes.html", struct {
		Usuarios  []modelos.Usuario
		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Pagina    string
		ASite     string
	}{
		Usuarios:  usuarios,
		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Nome:      nomeUsuario,
		Pagina:    "Usuarios",
		ASite:     aSite,
	})
}

//Carrega página principal
func CarregarPaginaFazerUpload(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios", config.APIURL)
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

	cookie, _ := cookies.Ler(r)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	busca := strings.ToLower(r.URL.Query().Get("a"))
	aSite, _ := cookie["asite"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "upload.html", struct {
		Usuarios  []modelos.Usuario
		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Busca     string
		Pagina    string
		ASite     string
	}{
		Usuarios:  usuarios,
		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Nome:      nomeUsuario,
		Busca:     busca,
		Pagina:    "Usuarios",
		ASite:     aSite,
	})
}

//Carrega página principal
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios", config.APIURL)
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

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "usuarios.html", struct {
		Usuarios                             []modelos.Usuario
		UsuarioID                            uint64
		Nome                                 string
		Login_NT                             string
		RE                                   string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{
		Usuarios:                             usuarios,
		UsuarioID:                            usuarioID,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		RE:                                   re,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
		Pagina:                               "Mapa Operações",
	})
}

//Carrega pagina de edição
func CarregarPaginaDeEdicaoDeUsuarios(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	usuariocID, erro := strconv.ParseUint(parametros["usuariocId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(usuariocID)
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuariocID)
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

	url2 := fmt.Sprintf("%s/sites", config.APIURL)
	response2, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url2, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response2.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var site []modelos.Site
	if erro = json.NewDecoder(response2.Body).Decode(&site); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]
	utils.ExecutarTemplate(w, "editar-usuario.html", struct {
		Usuario                              modelos.Usuario
		Site                                 []modelos.Site
		UsuarioID                            uint64
		Nome                                 string
		Login_NT                             string
		RE                                   string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{
		Usuario:                              usuario,
		Site:                                 site,
		UsuarioID:                            usuarioID,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		RE:                                   re,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
		Pagina:                               "Mapa Operações",
		//Pagina:    "Chamado: " + chamado.Chamado,
	})
}

//Carrega pagina de edição
func CarregarPaginaDeEdicaoDeSenha(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	usuariocID, erro := strconv.ParseUint(parametros["usuariocId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(usuariocID)
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuariocID)
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
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]
	fmt.Println(usuarioID)
	utils.ExecutarTemplate(w, "editar-senha-usuario.html", struct {
		Usuario                              modelos.Usuario
		UsuarioID                            uint64
		Nome                                 string
		Login_NT                             string
		RE                                   string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{
		Usuario:                              usuario,
		UsuarioID:                            usuarioID,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		RE:                                   re,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
		Pagina:                               "Mapa Operações",
		//Pagina:    "Chamado: " + chamado.Chamado,
	})
}

//Carrega página de layout
func CarregarPaginaLayout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	aSite, _ := cookie["asite"]
	nomeUsuario, _ := cookie["nome"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	perfilUsuario, _ := cookie["perfil"]
	fmt.Println(nomeUsuario)
	utils.ExecutarTemplate(w, "layout.html", struct {
		Pagina    string
		Nome      string
		UsuarioID uint64
		Perfil    string
		ASite     string
	}{
		Nome:      nomeUsuario,
		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Pagina:    "Layout",
		ASite:     aSite,
	})
}

//Carrega página de layout
func CarregarPaginaCPD(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	aSite, _ := cookie["asite"]
	nomeUsuario, _ := cookie["nome"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	perfilUsuario, _ := cookie["perfil"]
	fmt.Println(nomeUsuario)
	utils.ExecutarTemplate(w, "cpd.html", struct {
		Pagina    string
		Nome      string
		UsuarioID uint64
		Perfil    string
		ASite     string
	}{
		Nome:      nomeUsuario,
		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Pagina:    "CPD",
		ASite:     aSite,
	})
}

//Carrega página de layout
func CarregarPaginaMapaField(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]

	if v_imdb == "S" {

		utils.ExecutarTemplate(w, "mapa-field.html", struct {
			UsuarioID                            uint64
			Nome                                 string
			Login_NT                             string
			RE                                   string
			Cargo                                string
			V_HOMEOFFICE                         string
			V_HOMEOFFICE_CHAMADOS                string
			V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
			V_USUARIOS                           string
			V_IMPRESSORAS                        string
			V_BDC_POSTS                          string
			V_BDC_ADM                            string
			V_IMDB                               string
			V_GSA                                string
			V_CATRACA                            string
			V_BH                                 string
			V_MAPA_OPERACIONAL                   string
			SiteNome                             string
			Pagina                               string
		}{
			UsuarioID:                            usuarioID,
			Nome:                                 nomeUsuario,
			Login_NT:                             login_nt,
			RE:                                   re,
			Cargo:                                cargo,
			V_HOMEOFFICE:                         v_homeoffice,
			V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
			V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
			V_USUARIOS:                           v_usuarios,
			V_IMPRESSORAS:                        v_impressoras,
			V_BDC_POSTS:                          v_bdc_posts,
			V_BDC_ADM:                            v_bdc_adm,
			V_IMDB:                               v_imdb,
			V_GSA:                                v_gsa,
			V_CATRACA:                            v_catraca,
			V_BH:                                 v_bh,
			V_MAPA_OPERACIONAL:                   v_mapa_operacional,
			SiteNome:                             site_nome,
			Pagina:                               "Página GSA",
		})
	} else {

		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
			Pagina    string
			Nome      string
			UsuarioID uint64
			Perfil    string
			ASite     string
		}{
			Nome:      nomeUsuario,
			UsuarioID: usuarioID,
			Pagina:    "Página Inicial",
		})
	}
}

//Carrega página principal
func CarregarPaginaDeChamados(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/chamados", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var chamados []modelos.Chamado
	if erro = json.NewDecoder(response.Body).Decode(&chamados); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "chamados.html", struct {
		Chamados                             []modelos.Chamado
		UsuarioID                            uint64
		Nome                                 string
		Login_NT                             string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{
		Chamados:                             chamados,
		UsuarioID:                            usuarioID,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
		Pagina:                               "Página Inicial",
	})
}

//Carrega página principal
func CarregarPaginaDeChamadosAbertos(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/chamados", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var chamados []modelos.Chamado
	if erro = json.NewDecoder(response.Body).Decode(&chamados); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "chamados-abertos.html", struct {
		Chamados                             []modelos.Chamado
		UsuarioID                            uint64
		Nome                                 string
		Login_NT                             string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{
		UsuarioID:                            usuarioID,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
	})
}

//Carrega página principal
func CarregarPaginaDeChamadosProcessando(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/chamados", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var chamados []modelos.Chamado
	if erro = json.NewDecoder(response.Body).Decode(&chamados); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	aSite, _ := cookie["asite"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "chamados-processando.html", struct {
		Chamados  []modelos.Chamado
		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Pagina    string
		ASite     string
	}{
		Chamados:  chamados,
		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Nome:      nomeUsuario,
		Pagina:    "Chamados Processando",
		ASite:     aSite,
	})
}

//Carrega página principal
func CarregarPaginaDeChamadosADM(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/chamados", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var chamados []modelos.Chamado

	if erro = json.NewDecoder(response.Body).Decode(&chamados); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	aSite, _ := cookie["asite"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "chamados-adm.html", struct {
		Chamados  []modelos.Chamado
		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Pagina    string
		ASite     string
	}{
		Chamados:  chamados,
		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Nome:      nomeUsuario,
		Pagina:    "Chamados Aguardando ADM",
		ASite:     aSite,
	})
}

//Carrega página principal
func CarregarPaginaDeChamadosOak(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/chamados", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var chamados []modelos.Chamado
	if erro = json.NewDecoder(response.Body).Decode(&chamados); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	aSite, _ := cookie["asite"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "chamados-oak.html", struct {
		Chamados  []modelos.Chamado
		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Pagina    string
		ASite     string
	}{
		Chamados:  chamados,
		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Nome:      nomeUsuario,
		Pagina:    "Chamados Aguardando Oakmont",
		ASite:     aSite,
	})
}

//Carrega página principal
func CarregarPaginaDeChamadosFechados(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/chamados", config.APIURL)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var chamados []modelos.Chamado
	if erro = json.NewDecoder(response.Body).Decode(&chamados); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	aSite, _ := cookie["asite"]
	fmt.Println(usuarioID)

	utils.ExecutarTemplate(w, "chamados-fechados.html", struct {
		Chamados []modelos.Chamado

		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Pagina    string
		ASite     string
	}{
		Chamados: chamados,

		UsuarioID: usuarioID,
		Perfil:    perfilUsuario,
		Nome:      nomeUsuario,
		Pagina:    "Chamados Fechados",
		ASite:     aSite,
	})
}

//Carrega tela do formulario
func CarregarTelaDoFormularioMapa(w http.ResponseWriter, r *http.Request) {

	url2 := fmt.Sprintf("%s/sites", config.APIURL)
	response2, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url2, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response2.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response2)
		return
	}
	var site []modelos.Site
	if erro = json.NewDecoder(response2.Body).Decode(&site); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url3 := fmt.Sprintf("%s/clientes", config.APIURL)
	response3, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url3, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response3.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response3)
		return
	}
	var cliente []modelos.Cliente
	if erro = json.NewDecoder(response3.Body).Decode(&cliente); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url4 := fmt.Sprintf("%s/dacs", config.APIURL)
	response4, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url4, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response3.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response4)
		return
	}
	var dac []modelos.Dac
	if erro = json.NewDecoder(response4.Body).Decode(&dac); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url5 := fmt.Sprintf("%s/dominios", config.APIURL)
	response5, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url5, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	if response3.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response4)
		return
	}
	var dominio []modelos.Dominio
	if erro = json.NewDecoder(response5.Body).Decode(&dominio); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	re, _ := cookie["re"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]

	utils.ExecutarTemplate(w, "formulario-mapa.html", struct {
		Site                                 []modelos.Site
		Cliente                              []modelos.Cliente
		Dac                                  []modelos.Dac
		Dominio                              []modelos.Dominio
		UsuarioID                            uint64
		Nome                                 string
		Login_NT                             string
		RE                                   string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{

		Site:                                 site,
		Cliente:                              cliente,
		Dac:                                  dac,
		Dominio:                              dominio,
		UsuarioID:                            usuarioID,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		RE:                                   re,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
		Pagina:                               "Página Inicial",
	})
}

//Carrega tela do formulario
func CarregarTelaDoFormulario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]

	utils.ExecutarTemplate(w, "formulario.html", struct {
		UsuarioID                            uint64
		Nome                                 string
		Login_NT                             string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{
		UsuarioID:                            usuarioID,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
		Pagina:                               "Página Inicial",
	})
}

//Carrega tela do formulario
func CarregarTelaDoFormulario2(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	login_nt, _ := cookie["login_nt"]
	cargo, _ := cookie["cargo"]
	v_homeoffice, _ := cookie["v_homeoffice"]
	v_homeoffice_chamados, _ := cookie["v_homeoffice_chamados"]
	v_homeoffice_chamados_mudar_analista, _ := cookie["v_homeoffice_chamados_mudar_analista"]
	v_usuarios, _ := cookie["v_usuarios"]
	v_impressoras, _ := cookie["v_impressoras"]
	v_bdc_posts, _ := cookie["v_bdc_posts"]
	v_bdc_adm, _ := cookie["v_bdc_adm"]
	v_imdb, _ := cookie["v_imdb"]
	v_gsa, _ := cookie["v_gsa"]
	v_catraca, _ := cookie["v_catraca"]
	v_bh, _ := cookie["v_bh"]
	v_mapa_operacional, _ := cookie["v_mapa_operacional"]
	site_nome, _ := cookie["Site"]

	utils.ExecutarTemplate(w, "formulario2.html", struct {
		UsuarioID                            uint64
		Nome                                 string
		Login_NT                             string
		Cargo                                string
		V_HOMEOFFICE                         string
		V_HOMEOFFICE_CHAMADOS                string
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA string
		V_USUARIOS                           string
		V_IMPRESSORAS                        string
		V_BDC_POSTS                          string
		V_BDC_ADM                            string
		V_IMDB                               string
		V_GSA                                string
		V_CATRACA                            string
		V_BH                                 string
		V_MAPA_OPERACIONAL                   string
		SiteNome                             string
		Pagina                               string
	}{
		UsuarioID:                            usuarioID,
		Nome:                                 nomeUsuario,
		Login_NT:                             login_nt,
		Cargo:                                cargo,
		V_HOMEOFFICE:                         v_homeoffice,
		V_HOMEOFFICE_CHAMADOS:                v_homeoffice_chamados,
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: v_homeoffice_chamados_mudar_analista,
		V_USUARIOS:                           v_usuarios,
		V_IMPRESSORAS:                        v_impressoras,
		V_BDC_POSTS:                          v_bdc_posts,
		V_BDC_ADM:                            v_bdc_adm,
		V_IMDB:                               v_imdb,
		V_GSA:                                v_gsa,
		V_CATRACA:                            v_catraca,
		V_BH:                                 v_bh,
		V_MAPA_OPERACIONAL:                   v_mapa_operacional,
		SiteNome:                             site_nome,
		Pagina:                               "Página Inicial",
	})
}

//Carrega tela do formulario
func CarregarTelaDoFormularioTermo(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/chamados", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var chamados []modelos.Chamado
	if erro = json.NewDecoder(response.Body).Decode(&chamados); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	perfilUsuario, _ := cookie["perfil"]
	aSite, _ := cookie["asite"]

	utils.ExecutarTemplate(w, "formulario-termo.html", struct {
		Chamados  []modelos.Chamado
		UsuarioID uint64
		LoginNT   string
		Perfil    string
		Nome      string
		Pagina    string
		ASite     string
	}{
		Chamados:  chamados,
		UsuarioID: usuarioID,
		Nome:      nomeUsuario,
		Perfil:    perfilUsuario,
		Pagina:    "Formulário de Troca de Equipamentos",
		ASite:     aSite,
	})
}

//Carrega pagina de edição
func CarregarPaginaDoTermo(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	chamadoID, erro := strconv.ParseUint(parametros["chamadoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(chamadoID)
	url := fmt.Sprintf("%s/chamados/%d", config.APIURL, chamadoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	var chamado modelos.Chamado
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if erro = json.NewDecoder(response.Body).Decode(&chamado); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	nomeUsuario, _ := cookie["nome"]
	aSite, _ := cookie["asite"]
	utils.ExecutarTemplate(w, "termo.html", struct {
		Chamado   modelos.Chamado
		UsuarioID uint64
		LoginNT   string
		Nome      string
		Pagina    string
		ASite     string
	}{
		Chamado:   chamado,
		UsuarioID: usuarioID,
		Nome:      nomeUsuario,
		Pagina:    "Chamado: " + chamado.Chamado,
		ASite:     aSite,
	})
}
