package controllers

import (
	"app/src/config"
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
func CriarMapa(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	mapa, erro := json.Marshal(map[string]string{
		"operacao":          r.FormValue("operacao"),
		"vlan_dados":        r.FormValue("vlan_dados"),
		"vlan_voz":          r.FormValue("vlan_voz"),
		"config_contratual": r.FormValue("config_contratual"),
		"versao_windows":    r.FormValue("versao_windows"),
		"imagem":            r.FormValue("imagem"),
		"template":          r.FormValue("template"),
		"grupo_imdb":        r.FormValue("grupo_imdb"),
		"gravador":          r.FormValue("gravador"),
		"observacoes":       r.FormValue("observacoes"),
		"id_site":           r.FormValue("id_site"),
		"id_cliente":        r.FormValue("id_cliente"),
		"id_dac":            r.FormValue("id_dac"),
		"id_dominio":        r.FormValue("id_dominio"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(bytes.NewBuffer(mapa))

	url := fmt.Sprintf("%s/mapa/operacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(mapa))

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

//chama a api para cadastrar o usuario no db
func CriarChamado(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	chamado, erro := json.Marshal(map[string]string{
		"nome":                r.FormValue("nome"),
		"chamado":             r.FormValue("chamado"),
		"ativocpu":            r.FormValue("ativocpu"),
		"ativomonitor":        r.FormValue("ativomonitor"),
		"endereco":            r.FormValue("endereco"),
		"numero":              r.FormValue("numero"),
		"cep":                 r.FormValue("cep"),
		"senha":               r.FormValue("senha"),
		"transporte":          r.FormValue("transporte"),
		"acionamento":         r.FormValue("acionamento"),
		"status":              r.FormValue("status"),
		"bairro":              r.FormValue("bairro"),
		"obs":                 r.FormValue("obs"),
		"office":              r.FormValue("office"),
		"ramal":               r.FormValue("ramal"),
		"logindac":            r.FormValue("logindac"),
		"re":                  r.FormValue("re"),
		"ativoretornomonitor": r.FormValue("ativoretornomonitor"),
		"ativoretornocpu":     r.FormValue("ativoretornocpu"),
		"perifericomouse":     r.FormValue("perifericomouse"),
		"perifericoteclado":   r.FormValue("perifericoteclado"),
		"perifericohead":      r.FormValue("perifericohead"),
		"perifericorede":      r.FormValue("perifericorede"),
		"analistafield":       r.FormValue("analistafield"),
		"gerenteoperador":     r.FormValue("gerenteoperador"),
		"asite":               r.FormValue("asite"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	fmt.Println(bytes.NewBuffer(chamado))

	url := fmt.Sprintf("%s/chamados", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(chamado))

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
func AtualizarChamado(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	chamadoID, erro := strconv.ParseUint(parametros["chamadoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	r.ParseForm()

	chamado, erro := json.Marshal(map[string]string{
		"nome":                r.FormValue("nome"),
		"chamado":             r.FormValue("chamado"),
		"ativocpu":            r.FormValue("ativocpu"),
		"ativomonitor":        r.FormValue("ativomonitor"),
		"endereco":            r.FormValue("endereco"),
		"numero":              r.FormValue("numero"),
		"cep":                 r.FormValue("cep"),
		"senha":               r.FormValue("senha"),
		"transporte":          r.FormValue("transporte"),
		"acionamento":         r.FormValue("acionamento"),
		"status":              r.FormValue("status"),
		"bairro":              r.FormValue("bairro"),
		"obs":                 r.FormValue("obs"),
		"office":              r.FormValue("office"),
		"ramal":               r.FormValue("ramal"),
		"logindac":            r.FormValue("logindac"),
		"re":                  r.FormValue("re"),
		"ativoretornomonitor": r.FormValue("ativoretornomonitor"),
		"ativoretornocpu":     r.FormValue("ativoretornocpu"),
		"periferico_mouse":    r.FormValue("periferico_mouse"),
		"periferico_teclado":  r.FormValue("periferico_teclado"),
		"periferico_head":     r.FormValue("periferico_head"),
		"periferico_rede":     r.FormValue("periferico_rede"),
		"analistafield":       r.FormValue("analistafield"),
		"gerenteoperador":     r.FormValue("gerenteoperador"),
		"asite":               r.FormValue("asite"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/chamados/%d", config.APIURL, chamadoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(chamado))
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
func AtualizarMapa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	mapaID, erro := strconv.ParseUint(parametros["mapaId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	r.ParseForm()

	mapa, erro := json.Marshal(map[string]string{
		"operacao":          r.FormValue("operacao"),
		"vlan_dados":        r.FormValue("vlan_dados"),
		"vlan_voz":          r.FormValue("vlan_voz"),
		"config_contratual": r.FormValue("config_contratual"),
		"versao_windows":    r.FormValue("versao_windows"),
		"imagem":            r.FormValue("imagem"),
		"template":          r.FormValue("template"),
		"grupo_imdb":        r.FormValue("grupo_imdb"),
		"gravador":          r.FormValue("gravador"),
		"observacoes":       r.FormValue("observacoes"),
		"id_site":           r.FormValue("id_site"),
		"id_cliente":        r.FormValue("id_cliente"),
		"id_dac":            r.FormValue("id_dac"),
		"id_dominio":        r.FormValue("id_dominio"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/mapa/operacoes/%d", config.APIURL, mapaID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(mapa))
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
