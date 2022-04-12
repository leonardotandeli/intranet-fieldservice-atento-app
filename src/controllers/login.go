package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/modelos"
	"app/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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
