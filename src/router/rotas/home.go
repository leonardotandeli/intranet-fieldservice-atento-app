package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotaHome = []Rota{

	{
		URI:                "/home",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaPrincipal,
		RequerAutenticacao: true,
	},
}
