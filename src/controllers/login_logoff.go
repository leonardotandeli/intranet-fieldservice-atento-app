package controllers

import (
	"app/src/config"
	"app/src/cookies"
	"app/src/modelos"
	"app/src/requisicoes"
	"app/src/respostas"
	"app/src/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//Carrega a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	utils.ExecutarTemplate(w, "login.html", nil)
}

//Utiliza o usuario e senha para autenticação
func FazerLogin(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"login_nt": r.FormValue("login_nt"),
		"senha":    r.FormValue("senha"),
	})

	if erro != nil {
		//respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		fmt.Println("falha")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)

	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		//respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		http.Redirect(w, r, "/login"+"?message=Algo deu errado! Retorne daqui alguns minutos...", http.StatusFound)
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		http.Redirect(w, r, "/login"+"?message=Algo deu errado! Verifique seu login e senha e tente novamente.", http.StatusFound)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})

		return
	}

	if dadosAutenticacao.STATUS == "INATIVO" {
		http.Redirect(w, r, "/login"+"?message=Acesso bloqueado. Procure seu superior!", http.StatusFound)
		return
	}

	if erro = cookies.Salvar(w, dadosAutenticacao.ID, dadosAutenticacao.Token, dadosAutenticacao.NOME, dadosAutenticacao.LOGIN_NT, dadosAutenticacao.RE, dadosAutenticacao.CARGO, dadosAutenticacao.V_HOMEOFFICE, dadosAutenticacao.V_HOMEOFFICE_CHAMADOS, dadosAutenticacao.V_HOMEOFFICE_CHAMADOS_MUDAR_ANALISTA, dadosAutenticacao.V_USUARIOS, dadosAutenticacao.V_IMPRESSORAS, dadosAutenticacao.V_BDC_POSTS, dadosAutenticacao.V_BDC_ADM, dadosAutenticacao.V_IMDB, dadosAutenticacao.V_GSA, dadosAutenticacao.V_CATRACA, dadosAutenticacao.V_BH, dadosAutenticacao.V_CEP, dadosAutenticacao.V_MAPA_OPERACIONAL, dadosAutenticacao.V_MAPA_OPERACIONAL_ADM, dadosAutenticacao.Site.NOME, dadosAutenticacao.STATUS); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	http.Redirect(w, r, "/home", http.StatusFound)

}

func Deslogar(w http.ResponseWriter, r *http.Request) {

	//função para inserir dados dos cookies armazenados durante o login
	cookiesLoad, _ := cookies.InserirDadosNaPagina(r)
	userID := cookiesLoad.UsuarioID

	url := fmt.Sprintf("%s/usuarios/deslogar/%d", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
