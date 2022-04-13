package cookies

import (
	"app/src/config"
	"net/http"
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
	dadosCodificados, erro := s.Encode("dadosCookie", dados)
	if erro != nil {
		return erro
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "dadosCookie",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   99999,
	})
	return nil
}

//Ler retorna valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dadosCookie")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	if erro = s.Decode("dadosCookie", cookie.Value, &valores); erro != nil {
		return nil, erro
	}
	return valores, nil
}
func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dadosCookie",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
