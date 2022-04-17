package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasBaseDeConhecimento = []Rota{
	{
		URI:                "/base",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaInicialBase,
		RequerAutenticacao: true,
	},
	{
		URI:                "/formulario/base",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/formulario/categoria",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCriarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaBuscaCatOuCliente,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/{postId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/{postId}/editar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDePublicacao,
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
	{
		URI:                "/base/editar/categoria/{catId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeCategorias,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/editar/cliente/{clienteId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeClientes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/upload",
		Metodo:             http.MethodPost,
		Funcao:             controllers.UploadFile,
		RequerAutenticacao: true,
	},
}
