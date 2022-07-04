package rotas

import (
	"app/src/controllers"
	"net/http"
)

//rotasConsultas define as páginas de consultas
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
		Funcao:             controllers.CarregarPaginaDeConsultaDeIMDB,
		RequerAutenticacao: true,
	},
}
