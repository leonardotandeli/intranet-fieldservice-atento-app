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
func CarregarPaginaMapaString(w http.ResponseWriter, r *http.Request) {
	strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	strSite := strings.ToLower(r.URL.Query().Get("site"))

	strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")
	strSiteNoSpace := strings.ReplaceAll(strSite, " ", "+")

	fmt.Println(strClienteNoSpace)
	url := fmt.Sprintf("%s/mapa/busca?cliente=%s&site=%s", config.APIURL, strClienteNoSpace, strSiteNoSpace)
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

		utils.ExecutarTemplate(w, "mapa-string.html", struct {
			MapaOperacional                      []modelos.MapaOperacional
			Site                                 []modelos.Site
			Cliente                              []modelos.Cliente
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
			MapaOperacional:                      MapaOperacional,
			Site:                                 site,
			Cliente:                              cliente,
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
func CarregarPaginaMapa(w http.ResponseWriter, r *http.Request) {
	strCliente := strings.ToLower(r.URL.Query().Get("cliente"))
	strSite := strings.ToLower(r.URL.Query().Get("site"))

	strClienteNoSpace := strings.ReplaceAll(strCliente, " ", "+")
	strSiteNoSpace := strings.ReplaceAll(strSite, " ", "+")

	fmt.Println(strClienteNoSpace)
	url := fmt.Sprintf("%s/mapa/operacoes?cliente=%s&site=%s", config.APIURL, strClienteNoSpace, strSiteNoSpace)
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

		utils.ExecutarTemplate(w, "mapa.html", struct {
			MapaOperacional                      []modelos.MapaOperacional
			Site                                 []modelos.Site
			Cliente                              []modelos.Cliente
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
			MapaOperacional:                      MapaOperacional,
			Site:                                 site,
			Cliente:                              cliente,
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

//Carrega pagina de edição
func CarregarPaginaDeEdicaoMapa(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	mapaID, erro := strconv.ParseUint(parametros["mapaId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(mapaID)
	url := fmt.Sprintf("%s/mapa/operacoes/%d", config.APIURL, mapaID)
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
	utils.ExecutarTemplate(w, "editar-mapa.html", struct {
		MapaOperacional                      modelos.MapaOperacional
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
		MapaOperacional:                      MapaOperacional,
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
		Pagina:                               "Mapa Operações",
	})
}

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
