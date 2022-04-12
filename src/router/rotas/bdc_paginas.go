package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasBdcPaginas = []Rota{
	{
		URI:                "/base",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaInicialBase,
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
	},/*
		{
			URI:                "/base/novo",
			Metodo:             http.MethodGet,
			Funcao:             controllers.CarregarTelaDeCriarPublicacao,
			RequerAutenticacao: true,
		},
	
		{
			URI:                "/base/{publicacaoId}",
			Metodo:             http.MethodGet,
			Funcao:             controllers.CarregarPaginaDePublicacao,
			RequerAutenticacao: true,
		},
		{
			URI:                "/buscar",
			Metodo:             http.MethodGet,
			Funcao:             controllers.CarregarPaginaDeBusca,
			RequerAutenticacao: true,
		},*/
}
