package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasConfiguracao = []Rota{

	{
		URI:                "/configuracoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeConfiguracoes,
		RequerAutenticacao: true,
	},
}
