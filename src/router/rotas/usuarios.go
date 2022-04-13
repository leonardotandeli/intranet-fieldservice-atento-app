package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{

	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuariocId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuariocId}/editar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuariocId}/senha",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeSenha,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuariocId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},

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
