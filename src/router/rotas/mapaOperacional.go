package rotas

import (
	"app/src/controllers"
	"net/http"
)

//rotasMapaOperacional define as rotas do mapa operacional
var rotasMapaOperacional = []Rota{
	{
		URI:                "/mapa/operacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/busca/operacoes",
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
		URI:                "/mapa/adicionar/operacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDoFormularioMapa,
		RequerAutenticacao: true,
	},

	{
		URI:                "/mapa/adicionar/operacoes",
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
}
