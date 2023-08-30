// Package netsrv реализует telnet сервер для приема поисковых запросов.
package webapp

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	api    *API
	router *mux.Router
}

var apiData ApiData

func Start(d ApiData) error {
	apiData = d

	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = &API{router: srv.router}
	srv.api.Endpoints()

	return http.ListenAndServe(":8081", srv.router)
}
