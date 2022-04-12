package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasFerramentas = []Rota{

	{
		URI:                "/consulta/ativo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeAtivo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/cep",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeCEP,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/catraca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeCatraca,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/imdb",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeIMDB,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/bh",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeBH,
		RequerAutenticacao: true,
	},
}
