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
	"github.com/szyhf/go-excel"
)

//CriarUsuario chama a api para cadastrar o usuario no db
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":                   r.FormValue("nome"),
		"login_nt":               r.FormValue("login_nt"),
		"re":                     r.FormValue("re"),
		"cargo":                  r.FormValue("cargo"),
		"email":                  r.FormValue("email"),
		"v_usuarios":             r.FormValue("v_usuarios"),
		"v_bdc_posts":            r.FormValue("v_bdc_posts"),
		"v_bdc_adm":              r.FormValue("v_bdc_adm"),
		"v_imdb":                 r.FormValue("v_imdb"),
		"v_gsa":                  r.FormValue("v_gsa"),
		"v_mapa_operacional":     r.FormValue("v_mapa_operacional"),
		"v_mapa_operacional_adm": r.FormValue("v_mapa_operacional_adm"),
		"status":                 r.FormValue("status"),
		"id_site":                r.FormValue("id_site"),
		"senha":                  r.FormValue("senha"),
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

//AtualizarUsuario chama a API para alterar os dados de um usuário
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
		"nome":                   r.FormValue("nome"),
		"login_nt":               r.FormValue("login_nt"),
		"re":                     r.FormValue("re"),
		"cargo":                  r.FormValue("cargo"),
		"email":                  r.FormValue("email"),
		"v_usuarios":             r.FormValue("v_usuarios"),
		"v_bdc_posts":            r.FormValue("v_bdc_posts"),
		"v_bdc_adm":              r.FormValue("v_bdc_adm"),
		"v_imdb":                 r.FormValue("v_imdb"),
		"v_gsa":                  r.FormValue("v_gsa"),
		"v_mapa_operacional":     r.FormValue("v_mapa_operacional"),
		"v_mapa_operacional_adm": r.FormValue("v_mapa_operacional_adm"),
		"status":                 r.FormValue("status"),
		"id_site":                r.FormValue("id_site"),
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

//AtualizarSenha chama a API para alterar a senha de um usuário
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

	url := fmt.Sprintf("%s/usuarios/senha/%d", config.APIURL, usuarioID)
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

//UploadFileExcel chama a api para cadastrar o usuario no db utilizando a planilha do excel
func UploadFileExcel(w http.ResponseWriter, r *http.Request) {

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

	// junta o horário no formato unixnano com o nome do arquivo
	var nameFile = fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), handler.Filename)

	// cria o arquivo.
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
	// inicia a conexão para leitura da planilha
	conexao := excel.NewConnecter()
	erro := conexao.Open(nameFile[2:]) // abre a planilha no caminho do upload
	if erro != nil {
		fmt.Println(erro)

	}
	defer conexao.Close()
	// define a leitura da sheet "users"
	excel, erro := conexao.NewReader("users")
	if erro != nil {
		fmt.Println(erro)
	}

	for excel.Next() {
		var usuarios modelos.Usuario

		erro := excel.Read(&usuarios)
		if erro != nil {

			fmt.Println(erro)
		}

		if usuarios.CARGO == "M2" {
			usuarios.V_USUARIOS = "S"
		} else {
			usuarios.V_USUARIOS = "N"
		}

		switch usuarios.ID_SITE {
		case "ITF":
			usuarios.ID_SITE = "2"
		case "MAT":
			usuarios.ID_SITE = "3"
		case "BHP":
			usuarios.ID_SITE = "4"
		case "CBL":
			usuarios.ID_SITE = "5"
		case "CPG":
			usuarios.ID_SITE = "6"
		case "CSA":
			usuarios.ID_SITE = "7"
		case "CDN":
			usuarios.ID_SITE = "8"
		case "DLC":
			usuarios.ID_SITE = "9"
		case "FSA":
			usuarios.ID_SITE = "10"
		case "GOI":
			usuarios.ID_SITE = "11"
		case "GRU":
			usuarios.ID_SITE = "12"
		case "LBD":
			usuarios.ID_SITE = "13"
		case "MDR":
			usuarios.ID_SITE = "14"
		case "NSP":
			usuarios.ID_SITE = "15"
		case "OLC":
			usuarios.ID_SITE = "16"
		case "RJ":
			usuarios.ID_SITE = "17"
		case "POA":
			usuarios.ID_SITE = "18"
		case "REP":
			usuarios.ID_SITE = "19"
		case "NS2":
			usuarios.ID_SITE = "20"
		case "STN":
			usuarios.ID_SITE = "21"
		case "SAE":
			usuarios.ID_SITE = "22"
		case "STO":
			usuarios.ID_SITE = "23"
		case "SBE":
			usuarios.ID_SITE = "24"
		case "SBT":
			usuarios.ID_SITE = "25"
		case "SBC":
			usuarios.ID_SITE = "26"
		case "SCS":
			usuarios.ID_SITE = "27"
		case "SJC":
			usuarios.ID_SITE = "28"
		case "TLP":
			usuarios.ID_SITE = "29"
		case "URU":
			usuarios.ID_SITE = "30"
		case "ZSU":
			usuarios.ID_SITE = "31"
		case "ZL":
			usuarios.ID_SITE = "32"
		}

		usuarios.V_MAPA_OPERACIONAL_ADM = "S"
		usuarios.V_BDC_ADM = "S"
		usuarios.V_BDC_POSTS = "S"
		usuarios.V_IMDB = "S"
		usuarios.V_GSA = "S"
		usuarios.V_MAPA_OPERACIONAL = "S"
		usuarios.STATUS = "PRIMEIRO_ACESSO"
		usuarios.SENHA = "atento@22"

		r.ParseForm()

		usuario, erro := json.Marshal(map[string]string{
			"nome":                   usuarios.NOME,
			"login_nt":               usuarios.LOGIN_NT,
			"re":                     usuarios.RE,
			"cargo":                  usuarios.CARGO,
			"email":                  usuarios.EMAIL,
			"v_usuarios":             usuarios.V_USUARIOS,
			"v_bdc_posts":            usuarios.V_BDC_POSTS,
			"v_bdc_adm":              usuarios.V_BDC_ADM,
			"v_imdb":                 usuarios.V_IMDB,
			"v_gsa":                  usuarios.V_GSA,
			"v_mapa_operacional":     usuarios.V_MAPA_OPERACIONAL,
			"v_mapa_operacional_adm": usuarios.V_MAPA_OPERACIONAL_ADM,
			"id_site":                usuarios.ID_SITE,
			"status":                 usuarios.STATUS,
			"senha":                  usuarios.SENHA,
		})

		if erro != nil {
			respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})

		}
		//escreve os dados no banco de dados
		url := fmt.Sprintf("%s/usuarios", config.APIURL)
		response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(usuario))
		if erro != nil {
			respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})

		}

		defer response.Body.Close()
		if response.StatusCode >= 400 {
			respostas.TratarStatusCodeDeErro(w, response)
		}

	}

	//	respostas.JSON(w, response.StatusCode, nil)

}
