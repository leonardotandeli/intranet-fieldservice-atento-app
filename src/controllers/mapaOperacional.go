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
	//	fmt.Println(bytes.NewBuffer(mapa))

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

// DeletarMapa chama a API para deletar uma operação
func DeletarMapa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	mapaID, erro := strconv.ParseUint(parametros["mapaId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/mapa/operacoes/%d", config.APIURL, mapaID)
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

//UploadFileExcel chama a api para cadastrar o usuario no db utilizando a planilha do excel
func UploadFileExcelMapa(w http.ResponseWriter, r *http.Request) {

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
	fmt.Println(nameFile[45:])
	if nameFile[45:] != ".xlsx" {
		fmt.Println("teste")
		http.Redirect(w, r, "/criar-mapa-massa"+"?message=Erro ao ler o arquivo enviado. Por favor verifique a planilha e tente novamente.", http.StatusFound)
		return
	}

	defer dst.Close()

	// Copia o arquivo enviado para a pasta
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//	fmt.Println(nameFile[1:])

	conexaoExcel := excel.NewConnecter()
	conexaoExcel.Open(nameFile[2:])
	defer conexaoExcel.Close()

	leitor, erro := conexaoExcel.NewReader("operacoes")
	if erro != nil {
		fmt.Println("erro ao ler usuarios")
	}
	defer leitor.Close()

	for leitor.Next() {
		var operacoes modelos.MapaOperacional
		// Read a row into a struct.
		leitor.Read(&operacoes)
		//	fmt.Println(usuarios)

		switch operacoes.ID_DAC {
		case "CLIENTE":
			operacoes.ID_DAC = "1"
		case "AVAYA CLOUD":
			operacoes.ID_DAC = "2"
		case "BLZ A":
			operacoes.ID_DAC = "3"
		case "BLZ B":
			operacoes.ID_DAC = "4"
		case "BLZ C":
			operacoes.ID_DAC = "5"
		case "CAB":
			operacoes.ID_DAC = "6"
		case "GOI":
			operacoes.ID_DAC = "7"
		case "PEN A":
			operacoes.ID_DAC = "8"
		case "PEN B":
			operacoes.ID_DAC = "9"
		case "POA":
			operacoes.ID_DAC = "10"
		case "PRD":
			operacoes.ID_DAC = "11"
		case "RP":
			operacoes.ID_DAC = "12"
		case "SBC":
			operacoes.ID_DAC = "13"
		case "SBE A":
			operacoes.ID_DAC = "14"
		case "SBE B":
			operacoes.ID_DAC = "15"
		case "SBE C":
			operacoes.ID_DAC = "16"
		case "SBE D":
			operacoes.ID_DAC = "17"
		case "SBE E":
			operacoes.ID_DAC = "18"
		case "SJC":
			operacoes.ID_DAC = "19"
		default:
			operacoes.ID_DAC = "1"
		}

		switch operacoes.ID_SITE {
		case "ITF":
			operacoes.ID_SITE = "2"
		case "MAT":
			operacoes.ID_SITE = "3"
		case "BHP":
			operacoes.ID_SITE = "4"
		case "CBL":
			operacoes.ID_SITE = "5"
		case "CPG":
			operacoes.ID_SITE = "6"
		case "CSA":
			operacoes.ID_SITE = "7"
		case "CDN":
			operacoes.ID_SITE = "8"
		case "DLC":
			operacoes.ID_SITE = "9"
		case "FSA":
			operacoes.ID_SITE = "10"
		case "GOI":
			operacoes.ID_SITE = "11"
		case "GRU":
			operacoes.ID_SITE = "12"
		case "LBD":
			operacoes.ID_SITE = "13"
		case "MDR":
			operacoes.ID_SITE = "14"
		case "NSP":
			operacoes.ID_SITE = "15"
		case "OLC":
			operacoes.ID_SITE = "16"
		case "RJ":
			operacoes.ID_SITE = "17"
		case "POA":
			operacoes.ID_SITE = "18"
		case "REP":
			operacoes.ID_SITE = "19"
		case "NS2":
			operacoes.ID_SITE = "20"
		case "STN":
			operacoes.ID_SITE = "21"
		case "SAE":
			operacoes.ID_SITE = "22"
		case "STO":
			operacoes.ID_SITE = "23"
		case "SBE":
			operacoes.ID_SITE = "24"
		case "SBT":
			operacoes.ID_SITE = "25"
		case "SBC":
			operacoes.ID_SITE = "26"
		case "SCS":
			operacoes.ID_SITE = "27"
		case "SJC":
			operacoes.ID_SITE = "28"
		case "TLP":
			operacoes.ID_SITE = "29"
		case "URU":
			operacoes.ID_SITE = "30"
		case "ZSU":
			operacoes.ID_SITE = "31"
		case "ZL":
			operacoes.ID_SITE = "32"
		default:
			operacoes.ID_SITE = "1"
		}

		switch operacoes.ID_CLIENTE {
		case "ALELO":
			operacoes.ID_CLIENTE = "2"
		case "APPLE":
			operacoes.ID_CLIENTE = "3"
		case "ASURION":
			operacoes.ID_CLIENTE = "4"
		case "C6 BANK":
			operacoes.ID_CLIENTE = "5"
		case "BANCO CARREFOUR":
			operacoes.ID_CLIENTE = "6"
		case "BANCO DO BRASIL":
			operacoes.ID_CLIENTE = "7"
		case "BANCO PAN":
			operacoes.ID_CLIENTE = "8"
		case "BANCO SAFRA":
			operacoes.ID_CLIENTE = "9"
		case "BANCO SOFISA":
			operacoes.ID_CLIENTE = "10"
		case "BEMATECH":
			operacoes.ID_CLIENTE = "11"
		case "BENEFICIO FACIL":
			operacoes.ID_CLIENTE = "12"
		case "BMB":
			operacoes.ID_CLIENTE = "13"
		case "BMG":
			operacoes.ID_CLIENTE = "14"
		case "BOTICARIO":
			operacoes.ID_CLIENTE = "15"
		case "BRADESCO":
			operacoes.ID_CLIENTE = "16"
		case "BURGER KING":
			operacoes.ID_CLIENTE = "17"
		case "CAIXA SEGURADORA":
			operacoes.ID_CLIENTE = "18"
		case "CATENO":
			operacoes.ID_CLIENTE = "19"
		case "CETELEM":
			operacoes.ID_CLIENTE = "20"
		case "CIELO":
			operacoes.ID_CLIENTE = "21"
		case "CLARO":
			operacoes.ID_CLIENTE = "22"
		case "CONECTCAR":
			operacoes.ID_CLIENTE = "23"
		case "CONSULTING HOUSE":
			operacoes.ID_CLIENTE = "24"
		case "CTF":
			operacoes.ID_CLIENTE = "25"
		case "DASA":
			operacoes.ID_CLIENTE = "26"
		case "DECOLAR":
			operacoes.ID_CLIENTE = "27"
		case "DEMAIS AREAS":
			operacoes.ID_CLIENTE = "28"
		case "DISNEY":
			operacoes.ID_CLIENTE = "29"
		case "EASYNVEST":
			operacoes.ID_CLIENTE = "30"
		case "EDP":
			operacoes.ID_CLIENTE = "31"
		case "EDITORA GLOBO":
			operacoes.ID_CLIENTE = "32"
		case "ENEL":
			operacoes.ID_CLIENTE = "33"
		case "FACEBOOK":
			operacoes.ID_CLIENTE = "34"
		case "FIAT":
			operacoes.ID_CLIENTE = "35"
		case "FIRST DATA":
			operacoes.ID_CLIENTE = "36"
		case "FORD BRASIL":
			operacoes.ID_CLIENTE = "37"
		case "GOOGLE BRASIL":
			operacoes.ID_CLIENTE = "38"
		case "GPA":
			operacoes.ID_CLIENTE = "39"
		case "GRUPO ALIANCA":
			operacoes.ID_CLIENTE = "40"
		case "HONG":
			operacoes.ID_CLIENTE = "42"
		case "HUAWEI":
			operacoes.ID_CLIENTE = "43"
		case "ICATU":
			operacoes.ID_CLIENTE = "44"
		case "IFOOD":
			operacoes.ID_CLIENTE = "45"
		case "INTERMEDICA":
			operacoes.ID_CLIENTE = "46"
		case "INTERODONTO":
			operacoes.ID_CLIENTE = "47"
		case "ITAU":
			operacoes.ID_CLIENTE = "48"
		case "KROTON":
			operacoes.ID_CLIENTE = "49"
		case "LENOVO":
			operacoes.ID_CLIENTE = "50"
		case "LIVELO":
			operacoes.ID_CLIENTE = "51"
		case "LOSANGO":
			operacoes.ID_CLIENTE = "52"
		case "MARISA":
			operacoes.ID_CLIENTE = "53"
		case "MERCADO PAGO":
			operacoes.ID_CLIENTE = "54"
		case "MOTOROLA":
			operacoes.ID_CLIENTE = "55"
		case "NESTLE":
			operacoes.ID_CLIENTE = "56"
		case "OI":
			operacoes.ID_CLIENTE = "57"
		case "PEUGEOT":
			operacoes.ID_CLIENTE = "58"
		case "PRAVALER":
			operacoes.ID_CLIENTE = "59"
		case "QUALICORP":
			operacoes.ID_CLIENTE = "60"
		case "QUINTO ANDAR":
			operacoes.ID_CLIENTE = "61"
		case "RENNER":
			operacoes.ID_CLIENTE = "62"
		case "RIOT GAMES":
			operacoes.ID_CLIENTE = "63"
		case "SAMSUNG":
			operacoes.ID_CLIENTE = "64"
		case "SANTANDER":
			operacoes.ID_CLIENTE = "65"
		case "SEM PARAR":
			operacoes.ID_CLIENTE = "66"
		case "SHELL":
			operacoes.ID_CLIENTE = "67"
		case "SODEXO":
			operacoes.ID_CLIENTE = "68"
		case "SONY":
			operacoes.ID_CLIENTE = "69"
		case "STELO":
			operacoes.ID_CLIENTE = "70"
		case "SUL AMERICA":
			operacoes.ID_CLIENTE = "71"
		case "TELEFONICA":
			operacoes.ID_CLIENTE = "72"
		case "TIM":
			operacoes.ID_CLIENTE = "73"
		case "UNIDAS":
			operacoes.ID_CLIENTE = "75"
		case "UNILEVER":
			operacoes.ID_CLIENTE = "76"
		case "UNIMED BH":
			operacoes.ID_CLIENTE = "77"
		case "UNIMED FLORIANOPOLIS":
			operacoes.ID_CLIENTE = "78"
		case "UNIMED RJ":
			operacoes.ID_CLIENTE = "79"
		case "VELOE":
			operacoes.ID_CLIENTE = "80"
		case "VIA VAREJO":
			operacoes.ID_CLIENTE = "81"
		case "VIVO":
			operacoes.ID_CLIENTE = "82"
		case "WHITE MARTINS BRASIL":
			operacoes.ID_CLIENTE = "83"
		case "XP INVESTIMENTOS":
			operacoes.ID_CLIENTE = "84"
		case "Fleury":
			operacoes.ID_CLIENTE = "86"
		case "Carrefour":
			operacoes.ID_CLIENTE = "87"
		case "Froneri":
			operacoes.ID_CLIENTE = "88"
		case "UNIMED":
			operacoes.ID_CLIENTE = "89"
		case "UNIMED PORTO ALEGRE":
			operacoes.ID_CLIENTE = "91"
		case "3M":
			operacoes.ID_CLIENTE = "92"
		case "JCA":
			operacoes.ID_CLIENTE = "93"
		case "CVC":
			operacoes.ID_CLIENTE = "94"
		case "PASA":
			operacoes.ID_CLIENTE = "95"
		case "FACILY":
			operacoes.ID_CLIENTE = "96"
		default:
			operacoes.ID_CLIENTE = "1"
		}
		switch operacoes.ID_DOMINIO {
		case "ATENTOBR":
			operacoes.ID_DOMINIO = "1"
		case "ACIELO":
			operacoes.ID_DOMINIO = "2"
		case "AFACEBOOKBR":
			operacoes.ID_DOMINIO = "3"
		case "APOIOCASASBAHIA":
			operacoes.ID_DOMINIO = "4"
		case "ICATU-ATENTO":
			operacoes.ID_DOMINIO = "5"
		case "ARENNERBR":
			operacoes.ID_DOMINIO = "6"
		case "ITAUATENTOBR":
			operacoes.ID_DOMINIO = "7"
		case "AINTERMEDICA":
			operacoes.ID_DOMINIO = "8"
		case "CLIENTEAPBR":
			operacoes.ID_DOMINIO = "9"
		case "RECOVERYBR":
			operacoes.ID_DOMINIO = "10"
		default:
			operacoes.ID_DOMINIO = "1"
		}
		operacao := operacoes.OPERACAO
		/////////////operacoes.OPERACAOz////////
		//	fmt.Println(operacao)

		mapa, erro := json.Marshal(map[string]string{
			"operacao":          operacao,
			"vlan_dados":        operacoes.VLAN_DADOS,
			"vlan_voz":          operacoes.VLAN_VOZ,
			"config_contratual": operacoes.CONFIG_CONTRATUAL,
			"versao_windows":    operacoes.VERSAO_WINDOWS,
			"imagem":            operacoes.IMAGEM,
			"template":          operacoes.TEMPLATE,
			"grupo_imdb":        operacoes.GRUPO_IMDB,
			"gravador":          operacoes.GRAVADOR,
			"observacoes":       operacoes.OBSERVACOES,
			"id_site":           operacoes.ID_SITE,
			"id_cliente":        operacoes.ID_CLIENTE,
			"id_dac":            operacoes.ID_DAC,
			"id_dominio":        operacoes.ID_DOMINIO,
		})
		if erro != nil {
			respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
			return
		}
		//	fmt.Println(bytes.NewBuffer(mapa))

		url := fmt.Sprintf("%s/mapa/operacoes", config.APIURL)
		response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(mapa))

		if erro != nil {
			respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
			return
		}
		defer response.Body.Close()

		//	fmt.Println(strconv.Itoa(response.StatusCode) + operacoes.OPERACAO)

	}

	http.Redirect(w, r, "/criar-mapa-massa"+"?message=Import realizado com sucesso! Por favor validar.", http.StatusFound)

}
