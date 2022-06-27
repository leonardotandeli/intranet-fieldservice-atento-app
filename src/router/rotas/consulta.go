package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasConsultas = []Rota{
	{
		URI:                "/consulta/ativo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeAtivo,
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
		URI:                "/consulta/ad",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaConsultaAD,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/adg",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ConsultaAD,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/lapsg",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ConsultaLAPS,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/laps",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaConsultaLAPS,
		RequerAutenticacao: true,
	},
}
