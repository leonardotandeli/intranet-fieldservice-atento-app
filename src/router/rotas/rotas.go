package rotas

import (
	"app/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotasHomeOffice...)
	rotas = append(rotas, rotaDeslogar)
	rotas = append(rotas, rotasBdc...)
	rotas = append(rotas, rotasBdcPaginas...)
	rotas = append(rotas, rotasConfiguracao...)
	rotas = append(rotas, rotasUpload...)
	rotas = append(rotas, rotasFerramentas...)
	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)

		} else {
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}

	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router

}
