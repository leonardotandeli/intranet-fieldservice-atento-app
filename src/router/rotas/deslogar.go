package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotaDeslogar = Rota{
	URI:                "/deslogar",
	Metodo:             http.MethodGet,
	Funcao:             controllers.Deslogar,
	RequerAutenticacao: true,
}
