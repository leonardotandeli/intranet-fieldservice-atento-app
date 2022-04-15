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

//Carrega página de usuario
func CarregarPaginaDeConsultaDeAtivo(w http.ResponseWriter, r *http.Request) {
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

		utils.ExecutarTemplate(w, "consultar-ativo.html", struct {
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

//Carrega página de usuario
func CarregarPaginaDeConsultaDeCEP(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(nomeUsuario)
	utils.ExecutarTemplate(w, "consultar-cep.html", struct {
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

//Carrega página de usuario
func CarregarPaginaDeConsultaDeCatraca(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(nomeUsuario)
	utils.ExecutarTemplate(w, "consultar-catraca.html", struct {
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

//Carrega página de usuario
func CarregarPaginaDeConsultaDeIMDB(w http.ResponseWriter, r *http.Request) {
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

		utils.ExecutarTemplate(w, "consultar-imdb.html", struct {
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

//Carrega página de usuario
func CarregarPaginaDeConsultaDeBH(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(nomeUsuario)
	utils.ExecutarTemplate(w, "consultar-bh.html", struct {
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

//Carrega página do mapa de operações
func CarregarPaginaDeContatosField(w http.ResponseWriter, r *http.Request) {

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

	if v_mapa_operacional == "S" {

		utils.ExecutarTemplate(w, "contatos-field.html", struct {
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
			Pagina:                               "Mapa Operações",
		})
	} else {

		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
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
			Pagina:                               "Acesso Negado",
		})
	}
}

//Carrega página do mapa de operações
func CarregarPaginaDeContatosTsystems(w http.ResponseWriter, r *http.Request) {

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

	if v_mapa_operacional == "S" {

		utils.ExecutarTemplate(w, "contatos-tsystems.html", struct {
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
			Pagina:                               "Mapa Operações",
		})
	} else {

		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
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
			Pagina:                               "Acesso Negado",
		})
	}
}

//Carrega página do mapa de operações
func CarregarPaginaDeSalasDeCrise(w http.ResponseWriter, r *http.Request) {

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

	if v_mapa_operacional == "S" {

		utils.ExecutarTemplate(w, "salas-crise.html", struct {
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
			Pagina:                               "Mapa Operações",
		})
	} else {

		utils.ExecutarTemplate(w, "acesso-negado.html", struct {
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
			Pagina:                               "Acesso Negado",
		})
	}
}
