package rotas

import (
	"app/src/controllers"
	"net/http"
)

//rotasCategoriasBaseDeConhecimento define as rotas das Categorias da base de conhecimento
var rotasCategoriasBaseDeConhecimento = []Rota{
	{
		URI:                "/base-de-conhecimento/categorias/adicionar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCriarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/categoria",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaBuscaCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/alterar/categoria/{catId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeCategorias,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/alterar/categoria/{catId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/alterar/categoria/{catId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/alterar/categoria",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCategoria,
		RequerAutenticacao: true,
	},
}
