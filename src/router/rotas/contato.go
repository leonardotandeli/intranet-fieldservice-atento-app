package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasContatos = []Rota{
	{
		URI:                "/contatos/field",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeContatosField,
		RequerAutenticacao: true,
	},
	{
		URI:                "/contatos/tsystems",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeContatosTsystems,
		RequerAutenticacao: true,
	},
}
