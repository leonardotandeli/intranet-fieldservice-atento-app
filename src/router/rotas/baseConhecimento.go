package rotas

import (
	"app/src/controllers"
	"net/http"
)

//rotasBaseDeConhecimento define as rotas da base de conhecimento
var rotasBaseDeConhecimento = []Rota{
	{
		URI:                "/base-de-conhecimento",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaInicialBase,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/adicionar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeCriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaBusca,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/{postId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/{postId}/alterar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/base-de-conhecimento/upload",
		Metodo:             http.MethodPost,
		Funcao:             controllers.UploadFile,
		RequerAutenticacao: true,
	},
}
