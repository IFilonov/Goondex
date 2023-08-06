package webapp

import (
	"net/http"

	"github.com/gorilla/mux"
)

// API предоставляет интерфейс программного взаимодействия.
type API struct {
	router *mux.Router
}

type ApiData struct {
	IndexDocs []byte
	Docs      []byte
}

func (api *API) Endpoints() {
	api.router.HandleFunc("/index", api.index).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/docs", api.docs).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(apiData.IndexDocs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) docs(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(apiData.Docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
