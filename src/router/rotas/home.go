package rotas

import (
	"app/src/controllers"
	"net/http"
)

//rotaHome define a rota da página inicial
var rotaHome = []Rota{

	{
		URI:                "/home",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaPrincipal,
		RequerAutenticacao: true,
	},
}
