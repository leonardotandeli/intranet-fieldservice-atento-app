package controllers

import (
	"app/src/cookies"
	"app/src/modelos"
	"app/src/utils"
	"net/http"
	"strconv"
)

//Carrega página inicial
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", modelos.PageCookies{
		UsuarioID:                            usuarioID,
		Nome:                                 cookie["nome"],
		Login_NT:                             cookie["login_nt"],
		RE:                                   cookie["re"],
		Cargo:                                cookie["cargo"],
		V_HOMEOFFICE:                         cookie["v_homeoffice"],
		V_HOMEOFFICE_CHAMADOS:                cookie["v_homeoffice_chamados"],
		V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA: cookie["v_homeoffice_chamados_mudar_analista"],
		V_USUARIOS:                           cookie["v_usuarios"],
		V_IMPRESSORAS:                        cookie["v_impressoras"],
		V_BDC_POSTS:                          cookie["v_bdc_posts"],
		V_BDC_ADM:                            cookie["v_bdc_adm"],
		V_IMDB:                               cookie["v_imdb"],
		V_GSA:                                cookie["v_gsa"],
		V_CATRACA:                            cookie["v_catraca"],
		V_BH:                                 cookie["v_bh"],
		V_MAPA_OPERACIONAL:                   cookie["v_mapa_operacional"],
		SiteNome:                             cookie["Site"],
		Pagina:                               "Página Inicial",
	})
}
