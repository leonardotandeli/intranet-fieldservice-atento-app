package cookies

import (
	"app/src/config"
	"app/src/modelos"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// configura a sessão cookies
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

//registra autenticacao
func Salvar(w http.ResponseWriter, ID, token, NOME, LOGIN_NT, RE, CARGO, V_HOMEOFFICE, V_HOMEOFFICE_CHAMADOS, V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA, V_USUARIOS, V_IMPRESSORAS, V_BDC_POSTS, V_BDC_ADM, V_IMDB, V_GSA, V_CATRACA, V_BH, V_CEP, V_MAPA_OPERACIONAL, Site string) error {
	dados := map[string]string{
		"id":                                   ID,
		"token":                                token,
		"nome":                                 NOME,
		"login_nt":                             LOGIN_NT,
		"re":                                   RE,
		"cargo":                                CARGO,
		"v_homeoffice":                         V_HOMEOFFICE,
		"v_homeoffice_chamados":                V_HOMEOFFICE_CHAMADOS,
		"v_homeoffice_chamados_mudar_analista": V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA,
		"v_usuarios":                           V_USUARIOS,
		"v_impressoras":                        V_IMPRESSORAS,
		"v_bdc_posts":                          V_BDC_POSTS,
		"v_bdc_adm":                            V_BDC_ADM,
		"v_imdb":                               V_IMDB,
		"v_gsa":                                V_GSA,
		"v_catraca":                            V_CATRACA,
		"v_bh":                                 V_BH,
		"v_cep":                                V_CEP,
		"v_mapa_operacional":                   V_MAPA_OPERACIONAL,
		"Site":                                 Site,
	}
	dadosCodificados, erro := s.Encode("IntraFieldCookie", dados)
	if erro != nil {
		return erro
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "IntraFieldCookie",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   43150, //definido para 12 horas
	})
	return nil
}

//Ler retorna valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("IntraFieldCookie")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	if erro = s.Decode("IntraFieldCookie", cookie.Value, &valores); erro != nil {
		return nil, erro
	}
	return valores, nil
}
func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "IntraFieldCookie",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}

//InserirDadosNaPagina faz a leitura e insere os dados armazenados na página através do struct PageCookies
func InserirDadosNaPagina(r *http.Request) (modelos.PageCookies, error) {

	cookie, _ := Ler(r)
	var c modelos.PageCookies

	c.UsuarioID, _ = strconv.ParseUint(cookie["id"], 10, 64)
	c.Nome = cookie["nome"]
	c.Login_NT = cookie["login_nt"]
	c.RE = cookie["re"]
	c.Cargo = cookie["cargo"]
	c.V_HOMEOFFICE = cookie["v_homeoffice"]
	c.V_HOMEOFFICE_CHAMADOS = cookie["v_homeoffice_chamados"]
	c.V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA = cookie["v_homeoffice_chamados_mudar_analista"]
	c.V_USUARIOS = cookie["v_usuarios"]
	c.V_IMPRESSORAS = cookie["v_impressoras"]
	c.V_BDC_POSTS = cookie["v_bdc_posts"]
	c.V_BDC_ADM = cookie["v_bdc_adm"]
	c.V_IMDB = cookie["v_imdb"]
	c.V_GSA = cookie["v_gsa"]
	c.V_CATRACA = cookie["v_catraca"]
	c.V_BH = cookie["v_bh"]
	c.V_MAPA_OPERACIONAL = cookie["v_mapa_operacional"]
	c.SiteNome = cookie["Site"]

	return c, nil
}
