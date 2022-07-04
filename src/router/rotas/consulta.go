package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasConsultas = []Rota{
	{
		URI:                "/consulta/ativo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaGSA,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/imdb",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaIMDB,
		RequerAutenticacao: true,
	},
}
