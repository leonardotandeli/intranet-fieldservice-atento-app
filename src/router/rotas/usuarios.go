package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{

	{
		URI:                "/usuarios/novo",
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
		URI:                "/usuarios/{usuariocId}/alterar",
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
		URI:                "/usuarios/massa",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuariosExcel,
		RequerAutenticacao: true,
	},
	{
		URI:                "/upload/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.UploadFileExcel,
		RequerAutenticacao: true,
	},
}
