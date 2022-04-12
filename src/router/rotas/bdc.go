package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasBdc = []Rota{

	{
		URI:                "/base/novo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
