package rotas

import (
	"app/src/controllers"
	"net/http"
)

//define as rotas das páginas das salas de crise.
var rotasSalaDeCrise = []Rota{
	{
		URI:                "/salas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeSalasDeCrise,
		RequerAutenticacao: true,
	},
}
