package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/modelos"
	"app/src/requisicoes"
	"app/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//chama a api para cadastrar o usuario no db
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":               r.FormValue("nome"),
		"login_nt":           r.FormValue("login_nt"),
		"re":                 r.FormValue("re"),
		"cargo":              r.FormValue("cargo"),
		"email":              r.FormValue("email"),
		"v_usuarios":         r.FormValue("v_usuarios"),
		"v_bdc_posts":        r.FormValue("v_bdc_posts"),
		"v_bdc_adm":          r.FormValue("v_bdc_adm"),
		"v_imdb":             r.FormValue("v_imdb"),
		"v_gsa":              r.FormValue("v_gsa"),
		"v_mapa_operacional": r.FormValue("v_mapa_operacional"),
		"id_site":            r.FormValue("id_site"),
		"senha":              r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(bytes.NewBuffer(usuario))
	url := fmt.Sprintf("%s/usuarios", config.APIURL)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(usuario))

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// AtualizarPublicacao chama a API para atualizar uma publicação
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuariocId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	println(usuarioID)
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":               r.FormValue("nome"),
		"login_nt":           r.FormValue("login_nt"),
		"re":                 r.FormValue("re"),
		"cargo":              r.FormValue("cargo"),
		"email":              r.FormValue("email"),
		"v_usuarios":         r.FormValue("v_usuarios"),
		"v_bdc_posts":        r.FormValue("v_bdc_posts"),
		"v_bdc_adm":          r.FormValue("v_bdc_adm"),
		"v_imdb":             r.FormValue("v_imdb"),
		"v_gsa":              r.FormValue("v_gsa"),
		"v_mapa_operacional": r.FormValue("v_mapa_operacional"),
		"id_site":            r.FormValue("id_site"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 405 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

// AtualizarPublicacao chama a API para atualizar uma publicação
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuariocId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nova": r.FormValue("nova"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/atualizar-senha/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 405 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

//Utiliza o usuario e senha para autenticação
func FazerLogin(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"login_nt": r.FormValue("login_nt"),
		"senha":    r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)

	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, dadosAutenticacao.ID, dadosAutenticacao.Token, dadosAutenticacao.NOME, dadosAutenticacao.LOGIN_NT, dadosAutenticacao.RE, dadosAutenticacao.CARGO, dadosAutenticacao.V_HOMEOFFICE, dadosAutenticacao.V_HOMEOFFICE_CHAMADOS, dadosAutenticacao.V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA, dadosAutenticacao.V_USUARIOS, dadosAutenticacao.V_IMPRESSORAS, dadosAutenticacao.V_BDC_POSTS, dadosAutenticacao.V_BDC_ADM, dadosAutenticacao.V_IMDB, dadosAutenticacao.V_GSA, dadosAutenticacao.V_CATRACA, dadosAutenticacao.V_BH, dadosAutenticacao.V_CEP, dadosAutenticacao.V_MAPA_OPERACIONAL, dadosAutenticacao.Site.NOME); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	respostas.JSON(w, http.StatusOK, nil)

}

func Deslogar(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}
