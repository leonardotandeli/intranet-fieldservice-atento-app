package rotas

import (
	"app/src/controllers"
	"net/http"
)

//rotasClientesBaseDeConhecimento define as rotas de clientes na base de conhecimento
var rotasClientesBaseDeConhecimento = []Rota{
	{
		URI:                "/base-de-conhecimento/cliente/adicionar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCriarCliente,
		RequerAutenticacao: true,
	},

	{
		URI:                "/base-de-conhecimento/cliente",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaBuscaCatOuCliente,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/alterar/cliente/{clienteId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeClientes,
		RequerAutenticacao: true,
	},
}
