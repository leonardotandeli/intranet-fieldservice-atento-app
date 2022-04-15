package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasPaginas = []Rota{

	{
		URI:                "/configuracoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConfiguracoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/home",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaPrincipal,
		RequerAutenticacao: true,
	},

	{
		URI:                "/consulta/ativo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeAtivo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/cep",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeCEP,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/catraca",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeCatraca,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/imdb",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeIMDB,
		RequerAutenticacao: true,
	},
	{
		URI:                "/consulta/bh",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConsultaDeBH,
		RequerAutenticacao: true,
	},

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
		URI:                "/formulario/mapa",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarMapa,
		RequerAutenticacao: true,
	},

	{
		URI:                "/mapa/operacoes/{mapaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarMapa,
		RequerAutenticacao: true,
	},
}
