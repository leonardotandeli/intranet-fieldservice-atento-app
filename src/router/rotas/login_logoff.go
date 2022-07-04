package rotas

import (
	"app/src/controllers"
	"net/http"
)

//rotasLogin define as rotas de login e logoff
var rotasLogin = []Rota{
	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		Metodo:             http.MethodPost,
		Funcao:             controllers.FazerLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/deslogar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Deslogar,
		RequerAutenticacao: true,
	},
}
