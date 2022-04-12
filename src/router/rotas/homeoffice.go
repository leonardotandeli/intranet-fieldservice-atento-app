package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasHomeOffice = []Rota{

	{
		URI:                "/layout",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaLayout,
		RequerAutenticacao: true,
	},

	{
		URI:                "/cpd",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaCPD,
		RequerAutenticacao: true,
	},

	{
		URI:                "/mapa/operacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/busca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaMapaString,
		RequerAutenticacao: true,
	},

	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoMapa,
		RequerAutenticacao: true,
	},
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
	{
		URI:                "/salas-crise",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeSalasDeCrise,
		RequerAutenticacao: true,
	},
	{
		URI:                "/formulario/mapa",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDoFormularioMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/formulario/1",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDoFormulario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/formulario/2",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDoFormulario2,
		RequerAutenticacao: true,
	},

	{
		URI:                "/form-termo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDoFormularioTermo,
		RequerAutenticacao: true,
	},

	{
		URI:                "/formulario",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarChamado,
		RequerAutenticacao: true,
	},

	{
		URI:                "/formulario/mapa",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/chamados/{chamadoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarChamado,
		RequerAutenticacao: true,
	},
	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarMapa,
		RequerAutenticacao: true,
	},
	{
		URI:                "/chamados",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeChamados,
		RequerAutenticacao: true,
	},
	{
		URI:                "/chamados/abertos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeChamadosAbertos,
		RequerAutenticacao: true,
	},
	{
		URI:                "/chamados/processando",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeChamadosProcessando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/chamados/adm",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeChamadosADM,
		RequerAutenticacao: true,
	},
	{
		URI:                "/chamados/oakmont",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeChamadosOak,
		RequerAutenticacao: true,
	},
	{
		URI:                "/chamados/fechados",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeChamadosFechados,
		RequerAutenticacao: true,
	},
	{
		URI:                "/chamados/{chamadoId}/termo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDoTermo,
		RequerAutenticacao: true,
	},
}
