package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasSalaDeCrise = []Rota{
	{
		URI:                "/salas-crise",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeSalasDeCrise,
		RequerAutenticacao: true,
	},
}
