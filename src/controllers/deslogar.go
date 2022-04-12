package controllers

import (
	"app/src/cookies"
	"net/http"
)

func Deslogar(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}
