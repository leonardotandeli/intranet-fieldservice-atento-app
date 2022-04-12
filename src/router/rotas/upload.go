package rotas

import (
	"app/src/controllers"
	"net/http"
)

var rotasUpload = []Rota{
	{
		URI:                "/upload",
		Metodo:             http.MethodPost,
		Funcao:             controllers.FazerUpload,
		RequerAutenticacao: true,
	},
	{
		URI:                "/upload",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaFazerUpload,
		RequerAutenticacao: true,
	},
}
