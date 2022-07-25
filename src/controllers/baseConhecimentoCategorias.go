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

// AtualizarCategoria chama a API para atualizar uma categoria
func AtualizarCategoria(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	catID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	r.ParseForm()

	categoria, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/categorias/%d", config.APIURL, catID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(categoria))
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

// DeletarCategoria chama a API para deletar uma categoria
func DeletarCategoria(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	catID, erro := strconv.ParseUint(parametros["catId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/categorias/%d", config.APIURL, catID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
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

// AtualizarSubCategoria chama a API para atualizar uma categoria
func AtualizarSubCategoria(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	SubCatID, erro := strconv.ParseUint(parametros["subCatId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	r.ParseForm()

	subcategoria, erro := json.Marshal(map[string]string{
		"nome":         r.FormValue("nome"),
		"id_categoria": r.FormValue("id_categoria"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/subcategorias/%d", config.APIURL, SubCatID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(subcategoria))
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

// DeletarSubCategoria chama a API para deletar uma categoria
func DeletarSubCategoria(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	subCatID, erro := strconv.ParseUint(parametros["subCatId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/subcategorias/%d", config.APIURL, subCatID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
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

//chama a api para cadastrar uma categoria
func CriarCategoria(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"nome":       r.FormValue("nome"),
		"id_cliente": r.FormValue("id_cliente"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//	fmt.Println(bytes.NewBuffer(publicacao))

	url := fmt.Sprintf("%s/categorias", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))

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

//chama a api para cadastrar uma categoria
func CriarSubCategoria(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"nome":         r.FormValue("nome"),
		"id_categoria": r.FormValue("id_categoria"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//	fmt.Println(bytes.NewBuffer(publicacao))

	url := fmt.Sprintf("%s/subcategorias", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))

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
