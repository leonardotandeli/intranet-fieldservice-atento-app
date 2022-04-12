package controllers

import (
	"app/src/cookies"
	"app/src/utils"
	"fmt"
	"net/http"
	"strconv"
)

//Ca
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
