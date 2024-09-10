package routes

import (
	"net/http"
	handler "social-network/Handler"
)

func Routes(mux *http.ServeMux) {
	mux.HandleFunc("/", handler.Redirect)

	mux.HandleFunc("POST /login", handler.Login)
	mux.HandleFunc("POST /register", handler.Register)

	mux.HandleFunc("/home", handler.Home)
}
