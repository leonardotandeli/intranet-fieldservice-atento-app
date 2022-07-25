package controllers

import (
	"app/src/config"
	"app/src/modelos"
	"app/src/requisicoes"
	"app/src/respostas"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//CriarPublicacao chama a api para cadastrar uma nova publicação
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":          r.FormValue("titulo"),
		"conteudo":        r.FormValue("conteudo"),
		"id_categoria":    r.FormValue("id_categoria"),
		"id_subcategoria": r.FormValue("id_subcategoria"),
		"id_usuario":      r.FormValue("id_usuario"),
		"id_site":         r.FormValue("id_site"),
		"id_cliente":      r.FormValue("id_cliente"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//	fmt.Println(bytes.NewBuffer(publicacao))

	url := fmt.Sprintf("%s/base", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		//respostas.TratarStatusCodeDeErro(w, response)
		http.Redirect(w, r, "/base-de-conhecimento/"+"?message=Algo deu errado! Tente novamente mais tarde...", http.StatusFound)
		return
	}
	var post modelos.Post
	if erro = json.NewDecoder(response.Body).Decode(&post); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	http.Redirect(w, r, "/base-de-conhecimento/"+strconv.Itoa(int(post.IDPOST)), http.StatusFound)
	respostas.JSON(w, response.StatusCode, nil)

}

// AtualizarPublicacao chama a API para atualizar uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":          r.FormValue("titulo"),
		"conteudo":        r.FormValue("conteudo"),
		"id_categoria":    r.FormValue("id_categoria"),
		"id_subcategoria": r.FormValue("id_subcategoria"),
		"id_usuario":      r.FormValue("id_usuario"),
		"id_site":         r.FormValue("id_site"),
		"id_cliente":      r.FormValue("id_cliente"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/base/%d", config.APIURL, publicacaoID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(publicacao))
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

// DeletarPublicacao chama a API para deletar uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/base/%d", config.APIURL, publicacaoID)
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

//UploadFile recebe arquivos via upload
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// limita o upload para menos de 10mb
	r.ParseMultipartForm(10 << 20)

	// recebe o arquivo através do formulário
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Erro ao obter o arquivo.")
		fmt.Println(err)
		return
	}
	defer file.Close()

	//junta o horário em unixnano com o nome do arquivo
	var nameFile = fmt.Sprintf("./assets/uploads/%d%s", time.Now().UnixNano(), handler.Filename)

	//cria o arquivo.
	dst, err := os.Create(nameFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copia o arquivo enviado para a pasta
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//retorna uma mensagem em json para o tinymce omitindo o primeiro caracter [.]
	json.NewEncoder(w).Encode(map[string]string{"location": "" + nameFile[1:]})
}
