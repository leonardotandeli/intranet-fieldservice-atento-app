package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasMapaOperacional = []Rota{
	{
		URI:                "/mapa/operacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaMapaString,
		RequerAutenticacao: true,
	},

	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/formulario/mapa",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDoFormularioMapa,
		RequerAutenticacao: true,
	},

	{
		URI:                "/formulario/mapa",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarMapa,
		RequerAutenticacao: true,
	},

	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/criar-mapa-massa",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeOperacoesExcel,
		RequerAutenticacao: true,
	},

	{
		URI:                "/mapa/uploadExcel",
		Metodo:             http.MethodPost,
		Funcao:             controllers.UploadFileExcelMapa,
		RequerAutenticacao: true,
	},
}
