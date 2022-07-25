package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasBaseDeConhecimento = []Rota{
	{
		URI:                "/base-de-conhecimento",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaInicialBase,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/cliente/{clienteId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaBuscaCliente,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/posts",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaUltimosPosts,
		RequerAutenticacao: true,
	},
	///////////////////////////////////////////////////////////////////////////////
	{
		URI:                "/base-de-conhecimento/cliente/{clienteId}/categoria/{catId}/novo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/cliente/{clienteId}/categoria/novo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCriarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/cliente/{clienteId}/categoria/{catId}/subcategoria/novo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCriarSubCategoria,
		RequerAutenticacao: true,
	},
	/////////////////////////////////////////////////////////////////////////////////////////////
	{
		URI:                "/base-de-conhecimento/subcategorias",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarSubCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaBusca,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/cliente/{clienteId}/categoria/{catId}/subcategoria/{subCatId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaBuscaSubCategoria,
		RequerAutenticacao: true,
	},

	{
		URI:                "/base-de-conhecimento/cliente/{clienteId}/categoria/{catId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaBuscaCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/{postId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDePublicacao,
		RequerAutenticacao: true,
	},
	/////////////////////

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
		URI:                "/base/editar/categoria/{catId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/editar/categoria/{catId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/editar/categoria",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCategoria,
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

	{
		URI:                "/base/editar/subcategoria/{subCatId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarSubCategoria,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base/editar/subcategoria/{subCatId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarSubCategoria,
		RequerAutenticacao: true,
	},
}
